<template>
    <div>
        <div v-for="(post, index) in posts" :key="index" >
            <h1>{{post.title}}</h1>
            <p>{{post.text}}</p>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      posts: [
        {
          title: "СРОЧНЫЕ НОВОСТИ",
          text: "Обнаружен жж Артемия Лебедева!!!"
        },
        {
          title: "ВНЕЗАПНО",
          text: "Трава зелёная!!!"
        }
      ]
    };
  }
});
</script>
