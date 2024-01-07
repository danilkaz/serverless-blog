import viewPosts from "./components/blog/viewPosts.vue";
import addPost from "./components/blog/addPost.vue";
import NotFound from "./components/partials/notFound.vue";

export default [
  {
    path: "/",
    component: viewPosts,
  },
  {
    path: "/add",
    component: addPost,
  },
  {
    path: "/add/:id",
    component: addPost,
  },
  {
    path: "*",
    component: NotFound,
  },
];
