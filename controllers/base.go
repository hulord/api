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
	User models.User 	//登录的用户
}

const (
	KEY string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int = 600 // default 10 minutes
)


// JWT -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type MyCustomClaims struct {
	models.User
	jwt.StandardClaims
}

type JsonReturn struct {
	Status int		    `json:"status"`
	Message  string 	`json:"message"`
	Data interface{}	`json:"data"`		//Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`	
}

func (b *BaseController) Prepare() {
    //验证用户是否登录
	b.isLogin = false
	var LimitUri = []string{"/v1/artical/GetAll"}
	var isFlag  = false
	if utils.IsContain(LimitUri,b.Ctx.Request.RequestURI){
		 isFlag = true
	}
	if !isFlag && b.Ctx.Input.Header("Authorization") == "" {
		b.User = models.User{Username:"游客",Role:3} 			//Token为空是游客登录
		// beego.Error("without token, unauthorized !!")
		// b.ApiJsonReturn(1, "no permission","")
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
		fmt.Println( claims["username"].(string))
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

func GenerateToken(user models.User,expiredSeconds int) (tokenString string) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
   // Create the Claims
   mySigningKey := []byte(KEY)
   expireAt  := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
   //fmt.Println("token will be expired at ", time.Unix(expireAt, 0) )
   // pass parameter to this func or not
   user = models.User{Id:1,Username:"abc",Password:"ff",Gender:"ff",Age:"ff",Address:"ff",Email:"s",Role:1}
   claims := MyCustomClaims{
	   user,
	   jwt.StandardClaims{
		   ExpiresAt: expireAt,
		   Issuer:    user.Username,
		   IssuedAt:  time.Now().Unix(),
	   },
   }
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   tokenStr, err := token.SignedString(mySigningKey)
   if err != nil {
	   fmt.Println("generate json web token failed !! error :", err)
   }
   return tokenStr
}


// 校验token是否有效
func CheckToken(tokenString string) jwt.MapClaims {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return KEY, nil
    })
	if err != nil {
		fmt.Println("HS256的token解析错误，err:", err)
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("ParseHStoken:claims类型转换失败")
		return nil
	}
    return claims
}


//元素是否中数组中
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

