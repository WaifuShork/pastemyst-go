package pastemystgo

import "testing"

func TestUserExists(t *testing.T) {
	user, _ := UserExists("codemyst")
	if !user { 
		t.Errorf("Problem locating user.")
	}
}

func TestGetUser(t *testing.T) { 
	tests := []string{
		"codemyst",
		"waifushork",
	} 

	for _, tt := range tests { 
		user, _ := GetUser(tt)
		if !user.PublicProfile { 
			t.Errorf("Could not properly get user %s, please ensure their profile is public.", tt)
		}
	}
}