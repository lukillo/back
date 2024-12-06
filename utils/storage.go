package utils

import "sync"

var (
	emailList   = make(map[string]struct{})
	emailListMu sync.Mutex
)

func AddEmail(email string) {
	emailListMu.Lock()
	defer emailListMu.Unlock()
	emailList[email] = struct{}{}
}

func RemoveEmail(email string) {
	emailListMu.Lock()
	defer emailListMu.Unlock()
	delete(emailList, email)
}

func GetEmailList() map[string]struct{} {
	emailListMu.Lock()
	defer emailListMu.Unlock()
	return emailList
}
