package rediswatcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
	rds "github.com/redis/go-redis/v9"
)

type Watcher struct {
	l         sync.Mutex
	subClient rds.UniversalClient
	pubClient rds.UniversalClient
	options   WatcherOptions
	close     chan struct{}
	callback  func(string)
	ctx       context.Context
}

func DefaultUpdateCallback(e casbin.IEnforcer) func(string) {
	return func(msg string) {
		msgStruct := &MSG{}

		err := msgStruct.UnmarshalBinary([]byte(msg))
		if err != nil {
			log.Println(err)
			return
		}

		var res bool
		switch msgStruct.Method {
		case Update, UpdateForSavePolicy:
			err = e.LoadPolicy()
			res = true
		case UpdateForAddPolicy:
			res, err = e.SelfAddPolicy(msgStruct.Sec, msgStruct.Ptype, msgStruct.NewRule)
		case UpdateForAddPolicies:
			res, err = e.SelfAddPolicies(msgStruct.Sec, msgStruct.Ptype, msgStruct.NewRules)
		case UpdateForRemovePolicy:
			res, err = e.SelfRemovePolicy(msgStruct.Sec, msgStruct.Ptype, msgStruct.NewRule)
		case UpdateForRemoveFilteredPolicy:
			res, err = e.SelfRemoveFilteredPolicy(msgStruct.Sec, msgStruct.Ptype, msgStruct.FieldIndex, msgStruct.FieldValues...)
		case UpdateForRemovePolicies:
			res, err = e.SelfRemovePolicies(msgStruct.Sec, msgStruct.Ptype, msgStruct.NewRules)
		case UpdateForUpdatePolicy:
			res, err = e.SelfUpdatePolicy(msgStruct.Sec, msgStruct.Ptype, msgStruct.OldRule, msgStruct.NewRule)
		case UpdateForUpdatePolicies:
			res, err = e.SelfUpdatePolicies(msgStruct.Sec, msgStruct.Ptype, msgStruct.OldRules, msgStruct.NewRules)
		default:
			err = errors.New("unknown update type")
		}
		if err != nil {
			log.Println(err)
		}
		if !res {
			log.Println("callback update policy failed")
		}
	}
}

type MSG struct {
	Method      UpdateType
	ID          string
	Sec         string
	Ptype       string
	OldRule     []string
	OldRules    [][]string
	NewRule     []string
	NewRules    [][]string
	FieldIndex  int
	FieldValues []string
}

type UpdateType string

const (
	Update                        UpdateType = "Update"
	UpdateForAddPolicy            UpdateType = "UpdateForAddPolicy"
	UpdateForRemovePolicy         UpdateType = "UpdateForRemovePolicy"
	UpdateForRemoveFilteredPolicy UpdateType = "UpdateForRemoveFilteredPolicy"
	UpdateForSavePolicy           UpdateType = "UpdateForSavePolicy"
	UpdateForAddPolicies          UpdateType = "UpdateForAddPolicies"
	UpdateForRemovePolicies       UpdateType = "UpdateForRemovePolicies"
	UpdateForUpdatePolicy         UpdateType = "UpdateForUpdatePolicy"
	UpdateForUpdatePolicies       UpdateType = "UpdateForUpdatePolicies"
)

func (m *MSG) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary decodes the struct into a User
func (m *MSG) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}
	return nil
}

// NewWatcher creates a new Watcher to be used with a Casbin enforcer.
// `addr` is a Redis target string in the format "host:port".
// If you already have Redis clients initialized, you can pass them using WatcherOptions:
// - Use SubClient and PubClient for standalone Redis
//
// If clients are provided, `addr` can be an empty string ("").
// If clients are not provided, a new Redis client will be created using `addr`.
//
// Example with address:
//
//	w, err := rediswatcher.NewWatcher("127.0.0.1:6379", WatcherOptions{})
//
// Example with custom clients:
//
//	opts := WatcherOptions{
//	    SubClient: myRedisClient,
//	    PubClient: myRedisClient,
//	}
//	w, err := rediswatcher.NewWatcher("", opts)
func NewWatcher(addr string, option WatcherOptions) (persist.Watcher, error) {
	option.Options.Addr = addr
	initConfig(&option)
	w := &Watcher{
		ctx:   context.Background(),
		close: make(chan struct{}),
	}

	if err := w.initConfig(option); err != nil {
		return nil, err
	}

	if err := w.subClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}
	if err := w.pubClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}

	w.options = option

	w.subscribe()

	return w, nil
}

// NewWatcherWithCluster creates a new Watcher to be used with a Casbin enforcer.
// If SubClusterClient and PubClusterClient are already set in the WatcherOptions,
// the `addrs` string can be empty ("").
//
// addrs should be a comma-separated list of Redis cluster nodes in the format:
// "host1:port1,host2:port2,host3:port3"
//
// Example with address string:
//     w, err := rediswatcher.NewWatcherWithCluster("127.0.0.1:6379,127.0.0.1:6380,127.0.0.1:6381", WatcherOptions{})
//
// Example with pre-initialized clients:
//     opts := WatcherOptions{
//         SubClusterClient: existingSubClient,
//         PubClusterClient: existingPubClient,
//     }
//     w, err := rediswatcher.NewWatcherWithCluster("", opts)

func NewWatcherWithCluster(addrs string, option WatcherOptions) (persist.Watcher, error) {
	addrsStr := strings.Split(addrs, ",")
	option.ClusterOptions.Addrs = addrsStr
	initConfig(&option)
	var watcherSubClient rds.UniversalClient
	var watcherPubClient rds.UniversalClient

	watcherSubClient = option.SubClusterClient
	watcherPubClient = option.PubClusterClient

	if option.SubClusterClient == nil {
		watcherSubClient = rds.NewClusterClient(&rds.ClusterOptions{
			Addrs:    addrsStr,
			Username: option.ClusterOptions.Username,
			Password: option.ClusterOptions.Password,
		})
	}

	if option.PubClusterClient == nil {
		watcherPubClient = rds.NewClusterClient(&rds.ClusterOptions{
			Addrs:    addrsStr,
			Username: option.ClusterOptions.Username,
			Password: option.ClusterOptions.Password,
		})
	}

	w := &Watcher{
		subClient: watcherSubClient,
		pubClient: watcherPubClient,
		ctx:       context.Background(),
		close:     make(chan struct{}),
	}

	err := w.initConfig(option, true)
	if err != nil {
		return nil, err
	}

	if err := w.subClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}
	if err := w.pubClient.Ping(w.ctx).Err(); err != nil {
		return nil, err
	}

	w.options = option

	w.subscribe()

	return w, nil
}

func (wo *WatcherOptions) validate_wo() error {
	if wo.SubClient != nil && wo.SubClusterClient != nil {
		return fmt.Errorf("only one of SubClient or SubClusterClient should be set")
	}
	if wo.PubClient != nil && wo.PubClusterClient != nil {
		return fmt.Errorf("only one of SubClient or SubClusterClient should be set")
	}
	return nil
}

func (w *Watcher) initConfig(option WatcherOptions, cluster ...bool) error {
	var err error
	if option.OptionalUpdateCallback != nil {
		err = w.SetUpdateCallback(option.OptionalUpdateCallback)
	} else {
		err = w.SetUpdateCallback(func(string) {
			log.Println("Casbin Redis Watcher callback not set when an update was received")
		})
	}
	if err != nil {
		return err
	}

	err = option.validate_wo()
	if err != nil {
		return err
	}

	if option.SubClusterClient != nil {
		w.subClient = option.SubClusterClient
	} else if option.SubClient != nil {
		w.subClient = option.SubClient
	} else {
		// Neither provided, create new client
		if len(cluster) > 0 && cluster[0] {
			w.subClient = rds.NewClusterClient(&option.ClusterOptions)
		} else {
			w.subClient = rds.NewClient(&option.Options)
		}
	}

	if option.PubClusterClient != nil {
		w.pubClient = option.PubClusterClient
	} else if option.PubClient != nil {
		w.pubClient = option.PubClient
	} else {
		// Neither provided, create new client
		if len(cluster) > 0 && cluster[0] {
			w.pubClient = rds.NewClusterClient(&option.ClusterOptions)
		} else {
			w.pubClient = rds.NewClient(&option.Options)
		}
	}

	return nil
}

// NewPublishWatcher return a Watcher only publish but not subscribe
func NewPublishWatcher(addr string, option WatcherOptions) (persist.Watcher, error) {
	option.Options.Addr = addr
	w := &Watcher{
		pubClient: rds.NewClient(&option.Options),
		ctx:       context.Background(),
		close:     make(chan struct{}),
	}

	initConfig(&option)
	w.options = option

	return w, nil
}

// SetUpdateCallback sets the update callback function invoked by the watcher
// when the policy is updated. Defaults to Enforcer.LoadPolicy()
func (w *Watcher) SetUpdateCallback(callback func(string)) error {
	w.l.Lock()
	w.callback = callback
	w.l.Unlock()
	return nil
}

// Update publishes a message to all other casbin instances telling them to
// invoke their update callback
func (w *Watcher) Update() error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method: Update,
				ID:     w.options.LocalID,
			},
		).Err()
	})
}

// UpdateForAddPolicy calls the update callback of other instances to synchronize their policy.
// It is called after Enforcer.AddPolicy()
func (w *Watcher) UpdateForAddPolicy(sec, ptype string, params ...string) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method:  UpdateForAddPolicy,
				ID:      w.options.LocalID,
				Sec:     sec,
				Ptype:   ptype,
				NewRule: params,
			}).Err()
	})
}

// UpdateForRemovePolicy calls the update callback of other instances to synchronize their policy.
// It is called after Enforcer.RemovePolicy()
func (w *Watcher) UpdateForRemovePolicy(sec, ptype string, params ...string) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method:  UpdateForRemovePolicy,
				ID:      w.options.LocalID,
				Sec:     sec,
				Ptype:   ptype,
				NewRule: params,
			},
		).Err()
	})
}

// UpdateForRemoveFilteredPolicy calls the update callback of other instances to synchronize their policy.
// It is called after Enforcer.RemoveFilteredNamedGroupingPolicy()
func (w *Watcher) UpdateForRemoveFilteredPolicy(sec, ptype string, fieldIndex int, fieldValues ...string) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method:      UpdateForRemoveFilteredPolicy,
				ID:          w.options.LocalID,
				Sec:         sec,
				Ptype:       ptype,
				FieldIndex:  fieldIndex,
				FieldValues: fieldValues,
			},
		).Err()
	})
}

// UpdateForSavePolicy calls the update callback of other instances to synchronize their policy.
// It is called after Enforcer.RemoveFilteredNamedGroupingPolicy()
func (w *Watcher) UpdateForSavePolicy(model model.Model) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method: UpdateForSavePolicy,
				ID:     w.options.LocalID,
			},
		).Err()
	})
}

// UpdateForAddPolicies calls the update callback of other instances to synchronize their policies in batch.
// It is called after Enforcer.AddPolicies()
func (w *Watcher) UpdateForAddPolicies(sec string, ptype string, rules ...[]string) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method:   UpdateForAddPolicies,
				ID:       w.options.LocalID,
				Sec:      sec,
				Ptype:    ptype,
				NewRules: rules,
			},
		).Err()
	})
}

// UpdateForRemovePolicies calls the update callback of other instances to synchronize their policies in batch.
// It is called after Enforcer.RemovePolicies()
func (w *Watcher) UpdateForRemovePolicies(sec string, ptype string, rules ...[]string) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method:   UpdateForRemovePolicies,
				ID:       w.options.LocalID,
				Sec:      sec,
				Ptype:    ptype,
				NewRules: rules,
			},
		).Err()
	})
}

// UpdateForUpdatePolicy calls the update callback of other instances to synchronize their policy.
// It is called after Enforcer.UpdatePolicy()
func (w *Watcher) UpdateForUpdatePolicy(sec string, ptype string, oldRule, newRule []string) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method:  UpdateForUpdatePolicy,
				ID:      w.options.LocalID,
				Sec:     sec,
				Ptype:   ptype,
				OldRule: oldRule,
				NewRule: newRule,
			},
		).Err()
	})
}

// UpdateForUpdatePolicies calls the update callback of other instances to synchronize their policy.
// It is called after Enforcer.UpdatePolicies()
func (w *Watcher) UpdateForUpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	return w.logRecord(func() error {
		w.l.Lock()
		defer w.l.Unlock()
		return w.pubClient.Publish(
			context.Background(),
			w.options.Channel,
			&MSG{
				Method:   UpdateForUpdatePolicies,
				ID:       w.options.LocalID,
				Sec:      sec,
				Ptype:    ptype,
				OldRules: oldRules,
				NewRules: newRules,
			},
		).Err()
	})
}

func (w *Watcher) logRecord(f func() error) error {
	err := f()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (w *Watcher) unsubscribe(psc *rds.PubSub) error {
	return psc.Unsubscribe(w.ctx)
}

func (w *Watcher) subscribe() {
	w.l.Lock()
	sub := w.subClient.Subscribe(w.ctx, w.options.Channel)
	w.l.Unlock()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			err := sub.Close()
			if err != nil {
				log.Println(err)
			}
			err = w.pubClient.Close()
			if err != nil {
				log.Println(err)
			}
			err = w.subClient.Close()
			if err != nil {
				log.Println(err)
			}
		}()
		ch := sub.Channel()
		wg.Done()
		for msg := range ch {
			select {
			case <-w.close:
				return
			default:
			}
			data := msg.Payload
			msgStruct := &MSG{}
			err := msgStruct.UnmarshalBinary([]byte(data))
			if err != nil {
				log.Println(fmt.Printf("Failed to parse message: %s with error: %s\n", data, err.Error()))
			} else {
				isSelf := msgStruct.ID == w.options.LocalID
				if !(w.options.IgnoreSelf && isSelf) {
					w.callback(data)
				}
			}
		}
	}()
	wg.Wait()
}

func (w *Watcher) GetWatcherOptions() WatcherOptions {
	w.l.Lock()
	defer w.l.Unlock()
	return w.options
}

func (w *Watcher) Close() {
	w.l.Lock()
	defer w.l.Unlock()
	close(w.close)
	w.pubClient.Publish(w.ctx, w.options.Channel, "Close")
}
