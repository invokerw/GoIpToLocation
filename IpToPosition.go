package iptoposition

//根据string类型的IP地址来找到所在的地理位置，范范查找可以找到国家，国家ID
//稍微精确的查找可以找到所在城市，经纬度
//这里使用的是免费的文件GeoLite2_City.mmdb收费的查找的更加准确
//详细信息可以去github.com/oschwald/geoip2-golang查看也有两个文件的地址

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

//ret: country id  ,country name
func IpToCountryInfor(ipStr string) (uint, string) {

	db, err := geoip2.Open("./GeoLite2_City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipStr) //58.132.181.50
	//record, err := db.City(ip)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Name ID: %v\n", record.Country.GeoNameID)
	//fmt.Printf("English country name: %v\n", record.Country.NaIpmes["en"])
	//fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	//Output
	//Name ID: 1814991
	//English country name: China
	//ISO country code: CN
	return record.Country.GeoNameID, record.Country.Names["en"]
}
func IpToCityInfor(ipStr string) (string, float64, float64) {

	db, err := geoip2.Open("./GeoLite2_City.mmdb")
	//db, err := geoip2.Open("GeoIP2-City.mmdb")  //这个是收费的精度高
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipStr)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])  		葡萄牙城市名称
	//fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	//fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	//fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	//fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	//fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931
	return record.Subdivisions[0].Names["en"], record.Location.Latitude, record.Location.Longitude
}

/*
func main{
	id, strName := iptoposition.IpToCountryInfor("202.206.240.20")
	fmt.Println(id, strName)
	strCountryName, x, y := iptoposition.IpToCityInfor("202.206.240.20")
	fmt.Println(strCountryName, x, y)

	//Out put
	//1814991 China
	//Henan 34.6836 113.5325
}
*/
