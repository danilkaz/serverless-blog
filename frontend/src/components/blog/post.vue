<template>
    <div id="post-item">
        <h2 v-rand-color>{{ post.name }}</h2>
        <p>{{ post.text }}</p>
        <p style="text-align: right; font-size: x-small;">Дата создания: {{ new Date(Date.parse(post.created_at)).toLocaleDateString("ru") }} {{ new Date(Date.parse(post.created_at)).toLocaleTimeString("ru") }}</p>
    </div>
</template>
  
<script>
export default {
    name: 'viewPosts',
    props: {
        post: {
            type: Object,
            required: true
        }
    }
}
</script>
  
<style scoped>
#post-item {
    padding: 20px;
    margin: 20px 0;
    box-sizing: border-box;
    background: #eee;
}

.normal {
    text-decoration: none;
    color: inherit;
}
</style>
  