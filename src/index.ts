import Vue from "vue";
import PostComponent from './component/Post.vue'

let v = new Vue({
    el: "#app",
    template: `
    <div>
        <PostComponent />
    </div>`,
    components: {
        PostComponent
    }
});