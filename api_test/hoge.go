package main

import (
    "time"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "fmt"
    "os"
    "bufio"
)

const (
    URL     = "https://api.gnavi.co.jp/RestSearchAPI/v3/"
    CONSUMER_KEY = "e33d36e16e80413abfaf5caa66d49e66"
)

type AutoGenerated struct {
    Attributes struct {
        APIVersion string `json:"api_version"`
    } `json:"@attributes"`
    TotalHitCount int `json:"total_hit_count"`
    HitPerPage    int `json:"hit_per_page"`
    PageOffset    int `json:"page_offset"`
    Rest          []struct {
        Attributes struct {
            Order int `json:"order"`
        } `json:"@attributes"`
        ID         string    `json:"id"`
        UpdateDate time.Time `json:"update_date"`
        Name       string    `json:"name"`
        NameKana   string    `json:"name_kana"`
        Latitude   string    `json:"latitude"`
        Longitude  string    `json:"longitude"`
        Category   string    `json:"category"`
        URL        string    `json:"url"`
        URLMobile  string    `json:"url_mobile"`
        CouponURL  struct {
            Pc     string `json:"pc"`
            Mobile string `json:"mobile"`
        } `json:"coupon_url"`
        ImageURL struct {
            ShopImage1 string `json:"shop_image1"`
            ShopImage2 string `json:"shop_image2"`
            Qrcode     string `json:"qrcode"`
        } `json:"image_url"`
        Address  string `json:"address"`
        Tel      string `json:"tel"`
        TelSub   string `json:"tel_sub"`
        Fax      string `json:"fax"`
        Opentime string `json:"opentime"`
        Holiday  string `json:"holiday"`
        Access   struct {
            Line        string `json:"line"`
            Station     string `json:"station"`
            StationExit string `json:"station_exit"`
            Walk        string `json:"walk"`
            Note        string `json:"note"`
        } `json:"access"`
        ParkingLots string `json:"parking_lots"`
        Pr          struct {
            PrShort string `json:"pr_short"`
            PrLong  string `json:"pr_long"`
        } `json:"pr"`
        Code struct {
            Areacode      string   `json:"areacode"`
            Areaname      string   `json:"areaname"`
            Prefcode      string   `json:"prefcode"`
            Prefname      string   `json:"prefname"`
            AreacodeS     string   `json:"areacode_s"`
            AreanameS     string   `json:"areaname_s"`
            CategoryCodeL []string `json:"category_code_l"`
            CategoryNameL []string `json:"category_name_l"`
            CategoryCodeS []string `json:"category_code_s"`
            CategoryNameS []string `json:"category_name_s"`
        } `json:"code"`
        Budget     int    `json:"budget"`
        Party      int `json:"party int"`
        Lunch      int `json:"lunch int"`
        CreditCard string `json:"credit_card"`
        EMoney     string `json:"e_money"`
        Flags      struct {
            MobileSite   int `json:"mobile_site"`
            MobileCoupon int `json:"mobile_coupon"`
            PcCoupon     int `json:"pc_coupon"`
        } `json:"flags"`
    } `json:"rest"`
}

func getApi(text string) (result *http.Response, err error){
    values := url.Values{}
    values.Add("keyid", CONSUMER_KEY)
    get_url := URL + "?" + values.Encode() + "&" + "freeword=" + text
    resp, err := http.Get(get_url)
    return resp, err
}

func main() {
    // 標準入力
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Scan()
    text := stdin.Text()
    resp, err := getApi(text)

    if resp != nil {
        defer resp.Body.Close()
        var byteArray, _ = ioutil.ReadAll( resp.Body )
        fmt.Println( "ログイン成功!!!" )
        // fmt.Println( json_token )
        autogenerated := new(AutoGenerated)
        if err := json.Unmarshal(byteArray, &autogenerated); err != nil {
            log.Fatal(err)
        }
        fmt.Print(autogenerated.HitPerPage)
    } else {
        fmt.Println( err.Error() )
        return
    }

}

