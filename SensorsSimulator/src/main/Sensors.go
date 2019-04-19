package main

import (
	"log"
	"math"
	"time"
	"encoding/json"
)

type Sensor struct{
	Temperature float64
	Humidity uint8
	Timestamp int64
}

var sensor Sensor = Sensor{
	Temperature:50,
	Humidity:74,
	Timestamp:time.Now().Unix(),
}

//温度与湿度的对照表
var t2h = make(map[uint8]uint8)

func initSensors() {
	t2h[50] = 74
	t2h[49] = 74
	t2h[48] = 74
	t2h[47] = 73
	t2h[46] = 73
	t2h[45] = 73
	t2h[44] = 72
	t2h[43] = 72
	t2h[42] = 72
	t2h[41] = 71
	t2h[40] = 71
	t2h[39] = 70
	t2h[38] = 70
	t2h[37] = 69
	t2h[36] = 69
	t2h[35] = 68
	t2h[34] = 68
	t2h[33] = 67
	t2h[32] = 66
	t2h[31] = 66
	t2h[30] = 65
	t2h[29] = 64
	t2h[28] = 64
	t2h[30] = 65
	t2h[29] = 64
	t2h[28] = 64
	t2h[27] = 63
	t2h[26] = 62
	t2h[25] = 61
	t2h[24] = 60
	t2h[23] = 59
	t2h[22] = 58
	t2h[21] = 57
	t2h[20] = 56
	t2h[19] = 54
	t2h[18] = 53
	t2h[17] = 51
	t2h[16] = 50
	t2h[18] = 53
	t2h[17] = 51
	t2h[16] = 50
	t2h[15] = 48
	t2h[14] = 46
	t2h[13] = 45
	t2h[12] = 43
	t2h[11] = 40
	t2h[10] = 38
	t2h[9] = 36
	t2h[8] = 33
	t2h[7] = 31
	t2h[6] = 28
	t2h[5] = 24
	t2h[4] = 20
	t2h[3] = 16
	t2h[2] = 12
	t2h[1] = 8
	t2h[0] = 3
}

//将电梯当前状态转换为JSON格式数据
func ToJson() string {
	data, _ := json.Marshal(sensor)
	
	return string(data)
}

//开氏温度转换为摄氏温度
func K2T(k float64) float64 {

	return k - 273.15
}

//摄氏温度转换为开氏温度
func T2K(t float64) float64 {
	return 273.15 + t
}

//获取相对湿度
func SetRH(t float64) {
	
	t = math.Floor(t + 0.5)
	t = float64((int32(t * 1000) % 51000)/1000)
	
	b := uint8(t)
	log.Printf("b %d %d",b, t2h[b])
	
	sensor.Humidity = t2h[b]
}

//获取当前温度
func GetTemperature() float64 {
	return sensor.Temperature
}

//获取当前湿度
func GetHumidity() uint8 {
	return sensor.Humidity 
}

//运行传感器
func RunSensors(){
	for ;;{
		time.Sleep(time.Duration(1000)*time.Millisecond)
		Run()
	}	
}

func Run() {
	t := time.Now()

	hours := t.Hour()
	minutes := t.Minute()
	seconds := t.Second()
	
	switch {
	case hours > 0 && hours < 6:
		//从午夜00:00到凌晨6:00，温度应该处于下降，从15到5

		sensor.Temperature = float64(15 - 10*float64(hours*60*60+minutes*60+seconds)/float64(6*60*60))

	case hours >= 6 && hours < 14:
		//从6:00到14:00，温度应该处于上升，从5到30

		sensor.Temperature = float64(5 + 25*float64((hours-6)*60*60+minutes*60+seconds)/float64(8*60*60))

	case hours >= 14 && hours < 24:
		//从14:00到00:00，温度应该处于下降，从30到15

		sensor.Temperature = float64(30 - 15 * float64((hours-14)*60*60+minutes*60+seconds)/float64(10*60*60))
	}

	sensor.Timestamp = time.Now().Unix()
	SetRH(sensor.Temperature)
	
	log.Printf("Current Temperature %f", sensor.Temperature)
	log.Printf("Current Humidity %d", sensor.Humidity)
}
