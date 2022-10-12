
<template>
<view>
    <button @click="select">选择图片</button>
    <image mode="widthFix" :src="path"/>
    <ksp-cropper mode="free" :width="200" :height="140" :maxWidth="1024" :maxHeight="1024" :url="url" @cancel="oncancel" @ok="onok"></ksp-cropper>
</view>
</template>

<script>
    export default {
        data() {
            return {
                url: "",
                path: ""
            }
        },
        onLoad() {
        },
        methods: {
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
    }
</script>