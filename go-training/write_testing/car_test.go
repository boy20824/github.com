package car

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		t.Fatal("get errors : ", err)
	}

	if c == nil {
		t.Error("car should be nil")
	}

}

func TestNewWithAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)
}

//在vscode選取要測試的方法,並按下ctrl+shift+P之後選Go : Generate Unit Tests For Function
func TestCar_SetName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		c    *Car
		args args
		want string
	}{
		{name: "input name", //名稱描述
			c: &Car{
				Name:  "",
				Price: 0,
			},
			args: args{
				name: "Mark", //輸入的參數
			},
			want: "Mark"}, //預期的結果
		{name: "input name",
			c: &Car{
				Name:  "",
				Price: 0,
			},
			args: args{
				name: "barry",
			},
			want: "barry"},
	}
	for _, tt := range tests {
		tt := tt //宣告成區域變數防止被多執行序覆蓋
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() //開啟多條goroutine多執行
			log.Println(tt.args)
			if got := tt.c.SetName(tt.args.name); got != tt.want {
				t.Errorf("Car.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
