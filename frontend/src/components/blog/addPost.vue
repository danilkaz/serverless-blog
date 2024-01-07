<template>
    <div id="add-post">
        <h2>Добавить новый пост</h2>
        <form v-if="!submitted">
            <input type="text" v-model="post.name" required placeholder="Название">
            <textarea required v-model="post.text" placeholder="Текст"></textarea>
            <div class="buttons">
                <button class="btn btn-primary" v-on:click.prevent="submit">Создать</button>
                <button class="btn btn-danger" v-on:click.prevent="clear">Отменить</button>
            </div>
        </form>
        <div v-if="submitted" class="success">
            <h3>Пост успешно создан!</h3>
        </div>
    </div>
</template>
  
<script>
import axios from 'axios'
import getApiUrl from "../../apiUrl";
export default {
    name: 'addPost',
    data() {
        return {
            post: {
                name: '',
                text: '',
            },
            submitted: false,
        }
    }, methods: {
        submit() {
            if (this.post.name == '' || this.post.text == '') {
                return;
            }
            axios.post(`${getApiUrl()}/api/post`, this.post)
                .then(() => {
                    this.submitted = true;
                });
        }, clear() {
            this.post.name = '';
            this.post.text = '';
        }
    }
}
</script>
  
<style scoped>
#add-post {
    margin: 20px auto;
    max-width: 500px;
}

input,
textarea {
    display: block;
    width: 100%;
    padding: 8px;
    margin-bottom: 10px;
}

textarea {
    height: 150px;
}

.buttons {
    display: flex;
    justify-content: center;
}

.success {
    background-color: #4CAF50;
    padding: 10px;
    text-align: center;
}

.success h3 {
    margin: 0;
    color: #fff;
}

#preview-section {
    padding: 10px 20px;
    border: 1px dotted #ccc;
    margin: 30px 0;
}
</style>
  