<template>
    <nav>
        <p style="text-align: center;">Фронтенд: {{ this.frontendVersion }} Бэкенд: {{ this.backendVersion }} Реплика: {{ this.replicaId }}</p>
        <ul>
            <li><router-link to="/" exact>Все посты</router-link></li>
            <li><router-link to="/add" exact>Добавить новый пост</router-link></li>
        </ul>
    </nav>
</template>

<script>
import axios from 'axios'
import getApiUrl from "../../apiUrl";
export default {
    name: 'navbarApp',
    data() {
        return {
            frontendVersion: "",
            backendVersion: "",
            replicaId: "",
        }
    },
    mounted() {
        this.frontendVersion = require('../../../package.json').version;
        axios.get(`${getApiUrl()}/api/version`)
            .then(data => {
                this.backendVersion = data.data.version;
                this.replicaId = data.data.replica_id;
            });

    }
}
</script>

<style scoped>
ul {
    list-style-type: none;
    text-align: center;
    margin: 0;
}

li {
    display: inline-block;
    margin: 0 10px;
}

a {
    color: #fff;
    text-decoration: none;
    padding: 6px 8px;
    border-radius: 10px;
}

p {
    color: #fff;
    text-align: left;
    margin-top: 5px;
    /* padding: 14px; */
}

nav {
    background: #444;
    padding: 14px 0;
    margin-bottom: 40px;
}

.router-link-active {
    background: #eee;
    color: #444;
}
</style>