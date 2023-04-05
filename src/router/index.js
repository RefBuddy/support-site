// src/router/index.js
import { createRouter, createWebHistory } from "vue-router";
import HomePage from "../HomePage.vue";
import PrivacyPolicy from "../components/PrivacyPolicy.vue";

const routes = [
    {
        path: "/",
        name: "HomePage",
        component: HomePage,
    },
    {
        path: "/PrivacyPolicy",
        name: "PrivacyPolicy",
        component: PrivacyPolicy,
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
