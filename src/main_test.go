package main

// TBD
/*
func TestWenhookHandler(t *testing.T) {
	req := createTestRequest("POST", "/webhook/hoge", strings.NewReader(`{"message_id":31,"from":{"id":348776543,"is_bot":false,"first_name":"Kosuke(kosk1011)\ud83c\uddef\ud83c\uddf5","username":"kosk1011","language_code":"ja"},"chat":{"id":-358253261,"title":"test","type":"group","all_members_are_administrators":true},"date":1547731056,"text":"hello"}}`), t)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	wenhookHandler(res, req)
}

func TestErrHandler(t *testing.T) {
	req := createTestRequest("POST", "/hogehoge", strings.NewReader(`{"message_id":31,"from":{"id":348776543,"is_bot":false,"first_name":"Kosuke(kosk1011)\ud83c\uddef\ud83c\uddf5","username":"kosk1011","language_code":"ja"},"chat":{"id":-358253261,"title":"test","type":"group","all_members_are_administrators":true},"date":1547731056,"text":"hello"}}`), t)
	res := httptest.NewRecorder()
	errHandler(res, req)

}
*/
