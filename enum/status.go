package enum

const SHOW int = 1;
const HIDE int = 2;

var statusMap = map[int]string{SHOW:"显示",HIDE:"隐藏"};

const NOTKNOW string = "未知类型"

func GetStatus(status int) string {

	if (statusMap[status] == "") {
		return NOTKNOW;
	}
	return statusMap[status];

}