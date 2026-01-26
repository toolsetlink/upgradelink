# 生成patch
./hdiffz -f ./4.1.6.apk ./4.1.7.apk ./4.1.6-4.1.7.patch

# 合并patch
./hpatchz -f ./4.1.6.apk ./4.1.6-4.1.7.patch ./new-4.1.7.apk 