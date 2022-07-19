export default{
	//转换时间
	select() {
	    uni.chooseImage({
	        count: 1,
	        success: (rst) => {
	            // 设置url的值，显示控件
	            this.url = rst.tempFilePaths[0];
	        }
	    });
	},
	onok(ev) {
	    this.url = "";
	    this.path = ev.path;
	},
	oncancel() {
	    // url设置为空，隐藏控件
	    this.url = "";
	}
}