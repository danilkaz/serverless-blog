<template>
    <div class="container">
        <h2>Все посты</h2>
        <Loader v-if="isLoading" />
        <Post v-for="post in posts" :key="post.name" :post="post"/>
        <div v-if="!isLoading && posts.length == 0">
            <h3>Постов пока нет, добавьте первый!</h3>
        </div>
    </div>
</template>
  
<script>
import axios from 'axios'
import post from './post.vue';
import loader from '../partials/loader.vue';
import getApiUrl from "../../apiUrl";

export default {
    name: 'viewBlogs',
    components: {
        Post: post,
        Loader: loader
    },
    data() {
        return {
            posts: [],
            isLoading: true
        }
    },
    mounted() {
        this.isLoading = true;
        axios.get(`${getApiUrl()}/api/post`)
            .then(data => {
                this.posts = data.data.items;
                this.isLoading = false;
            });
    }

}
</script>
  