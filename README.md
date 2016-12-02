# GoIpToLocation
IP转地理位置的一个函数

两个函数
//ret: country id  ,country name
func IpToCountryInfor(ipStr string) (uint, string) 

//city name , 经纬度
func IpToCityInfor(ipStr string) (string, float64, float64) 
