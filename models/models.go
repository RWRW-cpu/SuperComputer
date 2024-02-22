package models

import "web/dao/mysql"

//
func LenTongDaoYanGan() (err error, list *[]IteaqLengTongDaoYanGan) {
	if err = mysql.DB.Find(&list).Error; err != nil {
		return err, nil
	}
	return nil, list
}

//环境空调温度
func GetAirConditioner1()(err error, list *[]IteaqAirConditioner1s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	
func GetAirConditioner2()(err error, list *[]IteaqAirConditioner2s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	
func GetAirConditioner3()(err error, list *[]IteaqAirConditioner3s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	
func GetAirConditioner4()(err error, list *[]IteaqAirConditioner4s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	
func GetAirConditioner5()(err error, list *[]IteaqAirConditioner5s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	

func GetAirConditioner6()(err error, list *[]IteaqAirConditioner6s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	
//继电器柱状图
func GetIteaqJiDianQiZhuZhuang()(err error, list *[]IteaqJiDianQiZhuZhuang){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	

//机房门禁
func GetIteaqMenJinJiFang()(err error, list *[]IteaqMenJinJiFang){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	

//门禁微模块
func GetIteaqMenJinWeiMoKuai()(err error, list *[]IteaqMenJinWeiMoKuai){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	

//其他电池
func GetIteaqOtherDianChi()(err error, list *[]IteaqOtherDianChi){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	


// 配电柜——精密
func GetIteaqPeiDianJingMi()(err error, list *[]IteaqPeiDianJingMi){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	

// 配电柜-配电1
func GetIteaqPeiDianPd1s()(err error, list *[]IteaqPeiDianPd1s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	
// 配电柜-配电1
func GetIteaqPeiDianUps1s()(err error, list *[]IteaqPeiDianUps1s){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	


// 冷通道温度
func GetIteaqLengTongDaoWenDu()(err error, list *[]IteaqTemp){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	


// UPS
func GetIteaqUps()(err error, list *[]IteaqUps){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	

// CPU
func GetGatewayCpuUsage()(err error, list *[]GatewayCpuUsage){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}	


// 硬盘
func GetGatewayDiskUsage()(err error, list *[]GatewayDiskUsage){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}

// 内存
func GetGatewayMemoryUsage()(err error, list *[]GatewayMemoryUsage){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}


// 登录
func GetGatewaySession()(err error, list *[]GatewaySession){
	if err = mysql.DB.Debug().Find(&list).Error; err != nil {
		return err, nil
	}
	return 
}