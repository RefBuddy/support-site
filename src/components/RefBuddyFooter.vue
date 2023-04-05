<template>
    <footer class="copyright">
        <p>
            &copy; {{ currentYear }} Ref Buddy. All rights reserved.
            <router-link to="/PrivacyPolicy">Privacy Policy</router-link>
        </p>
    </footer>
</template>

<script>
export default {
    name: "RefBuddyFooter",
    data() {
        return {
            currentYear: new Date().getFullYear(),
        };
    },
};
</script>

<style>
.copyright {
    text-align: center;
    font-size: small;
}
</style>
