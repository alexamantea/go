
// Package botusers contains struct with possible users
package botusers

//Structure Botuser
type Botuser struct {
	username string
	iconURL string

}
//New returns a botuser object
func New(userName string, iconURL string) Botuser{
	bot := Botuser{userName, iconURL}
	return bot
}
//users := []struct {
//	username, icon_url  string
//}{
//	{"Observer", "https://www.portallos.com.br/wp-content/uploads/2017/07/Observer-1170x680.jpg"},
//	{"Decorator", "https://s.hswstatic.com/gif/10-home-decorator-secrets-1.jpg"},
//	{"Barbie", "https://vignette.wikia.nocookie.net/boboiboy/images/4/4e/LargeAvatar_barbie.png"},
//}
//lAux := 0
//	for _, c := range users {
//		if lAux == pUserID{
//			return c.username, c.icon_url
//		}
//		lAux++
//	}
//	return "User Default",""
