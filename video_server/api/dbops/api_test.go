package dbops

import (
	"testing"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
	"strconv"
)

var(
	// Temp Vid to store parameter
	tempvid string
)

// init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// sub-tests for 4 tables operation
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("Ken Feng", "123456")
	if err != nil {
		t.Errorf("Error of Adding User: %s", err)
	}

}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("Ken Feng")
	if pwd != "123456" || err != nil {
		t.Errorf("Error of Getting User")
	}

}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("Ken Feng", "123456")
	if err != nil {
		t.Errorf("Error of Deleting User: %s", err)
	}

}

// check whether certain user information exists after delete it from table users
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("Ken Feng")
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error of Regetting: %s", err)
	}

	if pwd != "" {
		t.Errorf("Deleting User test Fail")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t * testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of Add VideoInfo: %s", err)
	}

	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of Get VideoInfo: %s", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of Delete VideoInfo: %s", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo After Deleting: %s", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "123456"
	aid := 1
	content := "This video is awful."

	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of Add Comments: %v\n",err)
	}
}

func testListComments(t *testing.T) {
	vid := "123456"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	_, err := ListComments(vid, from, to)
	// fmt.Printf("The result comments ammount is %d.\n", len(res))
	if err != nil {
		t.Errorf("Error of List Comments:%v\n", err)
	}
	/*
	for i, ele := range res {
		fmt.Printf("Comment: %d, %v \n", i, ele)
	}*/
}