package common

// 错误信息及docs文档
const (
	NewVersionMsg = "New version available"

	AlreadyLatestVersionMsg  = "Already latest version"
	AlreadyLatestVersionDocs = "It's already the latest version and no upgrade is required"

	Err1Msg  = "Internal service error"                                                                                                    // 无需修改，标准"服务内部错误"表述
	Err1Docs = "UpgradeLink internal service error. Please contact us for assistance: https://www.toolsetlink.com/upgrade/contact-us.html" // 优化：中文"服务内部错误"译为英文，保持文档说明语言统一（避免中英混杂）

	Err42901Msg  = "Request flow control restrictions"
	Err42901Docs = "The request flow control policy for the current upgrade task configuration has reached its limit. Please try again later."

	ErrApp1Msg  = "Request parameter error"
	ErrApp1Docs = ""
	ErrApp2Msg  = "The requested application does not exist"
	ErrApp2Docs = ""
	ErrApp3Msg  = "A usable version was not obtained"
	ErrApp3Docs = ""

	ErrFile1Msg  = "Request parameter error"
	ErrFile1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/file/get-file-upgrade-strategy.html"

	ErrFile2Msg  = "The requested application does not exist"
	ErrFile2Docs = "No application corresponding to the fileKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/file/app.html"

	ErrFile3Msg  = "No available version was found"
	ErrFile3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/file/app-version.html"

	ErrFile4Msg  = "Request parameter error"
	ErrFile4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/file/get-file-version.html"

	ErrFile5Msg  = "Request parameter error"
	ErrFile5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/file/get-file-download.html"

	ErrUrl1Msg  = "Request parameter error"
	ErrUrl1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/url/get-url-upgrade-strategy.html"

	ErrUrl2Msg  = "The requested application does not exist"
	ErrUrl2Docs = "No application corresponding to the urlKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/url/app.html"

	ErrUrl3Msg  = "No available version was found"
	ErrUrl3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/url/app-version.html"

	ErrUrl4Msg  = "Request parameter error"
	ErrUrl4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/url/get-url-version.html"

	ErrUrl5Msg  = "Request parameter error"
	ErrUrl5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/url/get-url-download.html"

	ErrConfiguration1Msg  = "Request parameter error"
	ErrConfiguration1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/configuration/get-configuration-upgrade-strategy.html"

	ErrConfiguration2Msg  = "The requested application does not exist"
	ErrConfiguration2Docs = "No application corresponding to the configurationKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/configuration/app.html"

	ErrConfiguration3Msg  = "No available version was found"
	ErrConfiguration3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/configuration/app-version.html"

	ErrConfiguration4Msg  = "Request parameter error"
	ErrConfiguration4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/configuration/get-configuration-version.html"

	ErrApk1Msg  = "Request parameter error"
	ErrApk1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/apk/get-apk-upgrade-strategy.html"

	ErrApk2Msg  = "The requested application does not exist"
	ErrApk2Docs = "No application corresponding to the apkKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/apk/app.html"

	ErrApk3Msg  = "No available version was found"
	ErrApk3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/apk/app-version.html"

	ErrApk4Msg  = "Request parameter error"
	ErrApk4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/apk/get-apk-version.html"

	ErrApk5Msg  = "Request parameter error"
	ErrApk5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/apk/get-apk-download.html"

	ErrTauri1Msg  = "Request parameter error"
	ErrTauri1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/tauri/get-tauri-upgrade-strategy.html"

	ErrTauri2Msg  = "The requested application does not exist"
	ErrTauri2Docs = "No application corresponding to the apkKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/tauri/app.html"

	ErrTauri3Msg  = "No available version was found"
	ErrTauri3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/tauri/app-version.html"

	ErrTauri4Msg  = "Request parameter error"
	ErrTauri4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/tauri/get-tauri-version.html"

	ErrTauri5Msg  = "Request parameter error"
	ErrTauri5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/tauri/get-tauri-download.html"

	ErrElectron1Msg  = "Request parameter error"
	ErrElectron1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/electron/get-electron-upgrade-strategy.html"

	ErrElectron2Msg  = "The requested application does not exist"
	ErrElectron2Docs = "No application corresponding to the apkKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/electron/app.html"

	ErrElectron3Msg  = "No available version was found"
	ErrElectron3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/electron/app-version.html"

	ErrElectron4Msg  = "Request parameter error"
	ErrElectron4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/electron/get-electron-version.html"

	ErrElectron5Msg  = "Request parameter error"
	ErrElectron5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/electron/get-electron-download.html"

	ErrMac1Msg  = "Request parameter error"
	ErrMac1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/mac/get-mac-upgrade-strategy.html"

	ErrMac2Msg  = "The requested application does not exist"
	ErrMac2Docs = "No application corresponding to the apkKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/mac/app.html"

	ErrMac3Msg  = "No available version was found"
	ErrMac3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/mac/app-version.html"

	ErrMac4Msg  = "Request parameter error"
	ErrMac4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/mac/get-mac-version.html"

	ErrMac5Msg  = "Request parameter error"
	ErrMac5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/mac/get-mac-download.html"

	ErrLnx1Msg  = "Request parameter error"
	ErrLnx1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/lnx/get-lnx-upgrade-strategy.html"

	ErrLnx2Msg  = "The requested application does not exist"
	ErrLnx2Docs = "No application corresponding to the apkKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/lnx/app.html"

	ErrLnx3Msg  = "No available version was found"
	ErrLnx3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/lnx/app-version.html"

	ErrLnx4Msg  = "Request parameter error"
	ErrLnx4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/lnx/get-lnx-version.html"

	ErrLnx5Msg  = "Request parameter error"
	ErrLnx5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/lnx/get-lnx-download.html"

	ErrWin1Msg  = "Request parameter error"
	ErrWin1Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/win/get-win-upgrade-strategy.html"

	ErrWin2Msg  = "The requested application does not exist"
	ErrWin2Docs = "No application corresponding to the apkKey was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/win/app.html"

	ErrWin3Msg  = "No available version was found"
	ErrWin3Docs = "No corresponding application version was found. Please confirm that it has been correctly created in the system. Refer to the documentation: https://www.toolsetlink.com/upgrade/recommend/win/app-version.html"

	ErrWin4Msg  = "Request parameter error"
	ErrWin4Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/win/get-win-version.html"

	ErrWin5Msg  = "Request parameter error"
	ErrWin5Docs = "Please refer to the documentation: https://www.toolsetlink.com/upgrade/api/win/get-win-download.html"
)
