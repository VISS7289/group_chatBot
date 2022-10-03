export default {
	//搜索防抖
	Debounce: function(fn, wait = 200) {
		let timer
		return function() {
			let context = this
			let args = arguments
			if (timer) clearTimeout(timer)
			timer = setTimeout(() => {
				fn.apply(context, args)
			}, wait)
		}
	}
}
