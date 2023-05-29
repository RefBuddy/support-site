<template>
    <div class="ref-buddy-page">
        <Header />
        <main class="main">
            <Introduction />
            <Features />
            <Screenshots />
            <MobileStores />
            <ContactForm />
            <RefBuddyFooter />
        </main>
    </div>
</template>

<script>
import Header from "./components/RefBuddyHeader.vue";
import Introduction from "./components/RefBuddyIntro.vue";
import Features from "./components/FeaturesList.vue";
import Screenshots from "./components/RefBuddyScreenshots.vue";
import MobileStores from "./components/MobileStoreButtons.vue";
import ContactForm from "./components/ContactForm.vue";
import RefBuddyFooter from "./components/RefBuddyFooter.vue";

export default {
    components: {
        Header,
        Introduction,
        Features,
        Screenshots,
        MobileStores,
        ContactForm,
        RefBuddyFooter,
    },
};
</script>

<style>
@import "@/css/CommonStyles.css";
</style>
