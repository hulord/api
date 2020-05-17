package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	utils "api/utils"
	"strings"
	"fmt"
	"api/models"
	"time" 	
)

type BaseController struct {
	beego.Controller
	isLogin bool     	//验证是否登录
	Role int
	Username string
}

const (
	KEY string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int = 600 // default 10 minutes
)

type JsonReturn struct {
	Status int		    `json:"status"`
	Message  string 	`json:"message"`
	Data interface{}	`json:"data"`		//Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`	
}

func (b *BaseController) Prepare() {
    //验证用户是否登录
	b.isLogin = false
	var LimitUri = []string{"/v1/artical/getall","/v1/artical/([1-9]+)","/v1/user/login"}
	var isFlag  = false
	if utils.IsContain(LimitUri,b.Ctx.Request.RequestURI){
		 isFlag = true
	}

	if  b.Ctx.Input.Header("Authorization") == "" {
		//允许游客查看的路由
		if isFlag {
			b.Role = 3
			b.Username = "游客"  //Token为空是游客登录
		//禁止路由
		}else{
			 beego.Error("without token, unauthorized !!")
			 b.ApiJsonReturn(1, "no permission","")
		}
	} 
	if !isFlag && b.Ctx.Input.Header("Authorization") != "" {
		authString := b.Ctx.Input.Header("Authorization")
		beego.Debug("AuthString:", authString)
		kv := strings.Split(authString, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			beego.Error("AuthString invalid:", authString)
			b.ApiJsonReturn(1, "AuthString invalid","")
		}
		// //检验Token是否成功
		claims,err := ParseToken(kv[1])
		
		if err!=nil {
			b.ApiJsonReturn(1, "AuthString invalid","")
		}
		b.isLogin = true
		b.Role = claims["role"].(int)
		b.Username = claims["username"].(string)
	}
}

func (b *BaseController) ApiJsonReturn(status int,message string,data interface{}) {
	var JsonReturn JsonReturn
	JsonReturn.Status = status
	JsonReturn.Message = message
	JsonReturn.Data = data
	b.Data["json"] = JsonReturn		//将结构体数组根据tag解析为json
	b.ServeJSON()					//对json进行序列化输出
	b.StopRun()						//终止执行逻辑
}


func CreateToken(user models.User,expiredSeconds int)(tokenss string,err error){
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
   // Create the Claims
   mySigningKey := []byte(KEY)
   expireAt  := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	//自定义claim
	claim := jwt.MapClaims{
		"username": user.Username,
		"Role":user.Role,
		"nbf":  time.Now().Unix(),
		"iat":  time.Now().Unix(),
		"exp":  expireAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenss,err  = token.SignedString(mySigningKey) 
	if err != nil {
		beego.Error("generate json web token failed !! error :", err)
	}
	return
}

func secret()jwt.Keyfunc{
    return func(token *jwt.Token) (interface{}, error) {
        return []byte(KEY),nil
    }
}


func ParseToken(tokenss string)(j jwt.MapClaims,err error){
    token,err := jwt.Parse(tokenss,secret())
    if err != nil{
        return nil,err
    }
    claim,ok := token.Claims.(jwt.MapClaims)
    if !ok{
		fmt.Println("111")

    }
    //验证token，如果token被修改过则为false
    if  !token.Valid{
	 fmt.Println("222")
    }

    return claim,nil
}

//元素是否中数组中
// func IsContain(items []string, item string)  {
// 	fmt.Println("abc")
// 	for _, eachItem := range items {
// 		compile := regexp.MustCompile(eachItem)
// 		submatch := compile.FindAllSubmatch(item, -1)
// 		fmt.Println(submatch)

// 	}
// 	return false
// }

