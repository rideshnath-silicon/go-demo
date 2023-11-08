package helper

func messages(MessageCode int) string {
	var messages string
	switch MessageCode {
	case 100:
		messages = "Data found successfully"
	case 101:
		messages = "Data not found "
	case 102:
		messages = "Data Inserted successfully"
	case 103:
		messages = "Data Updated Successfully"
	case 104:
		messages = "Data Delere Successfully"
	case 200:
		messages = "ID not found"
	case 201:
		messages = "Error In update"
	case 202:
		messages = "Error in create"

	default:
		messages = "okk"
	}
	return messages
}
