package main

import "testing"

func Test_stringbuilder(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringJia(tt.args.a); got != tt.want {
				t.Errorf("stringbuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStringJia(b *testing.B) {
	for i:=0;i<b.N;i++ {
		StringJia("SLEECT m.id as MessageId ,m.shop_id as ShopId,m.content as Content, m.message_type as MessageType,m.create_time as CreateTime ,mr.member_main as MemberMain,mr.is_read as IsRead")
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i:=0;i<b.N;i++ {
		StringBuilder("SLEECT m.id as MessageId ,m.shop_id as ShopId,m.content as Content, m.message_type as MessageType,m.create_time as CreateTime ,mr.member_main as MemberMain,mr.is_read as IsRead")
	}
}

func BenchmarkByteString(b *testing.B) {
	for i:=0;i<b.N;i++ {
		ByteString("SLEECT m.id as MessageId ,m.shop_id as ShopId,m.content as Content, m.message_type as MessageType,m.create_time as CreateTime ,mr.member_main as MemberMain,mr.is_read as IsRead")
	}
}