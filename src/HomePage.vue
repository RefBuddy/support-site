<template>
    <div class="ref-buddy-page">
        <Header />
        <main class="main">
            <Introduction />
            <MobileStores />
            <ContactForm />
            <RefBuddyFooter />
        </main>
    </div>
</template>

<script>
import Header from "./components/RefBuddyHeader.vue";
import Introduction from "./components/RefBuddyIntro.vue";
import MobileStores from "./components/MobileStoreButtons.vue";
import ContactForm from "./components/ContactForm.vue";
import RefBuddyFooter from "./components/RefBuddyFooter.vue";

export default {
    components: {
        Header,
        Introduction,
        MobileStores,
        ContactForm,
        RefBuddyFooter,
    },
};
</script>

<style>
@import "@/css/CommonStyles.css";
</style>
