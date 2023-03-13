<template>
    <div class="ref-buddy-page">
        <Header />
        <main class="main">
            <Introduction />
            <Features />
            <Screenshots />
            <ContactForm />
        </main>
    </div>
</template>

<script>
import Header from "./components/RefBuddyHeader.vue";
import Introduction from "./components/RefBuddyIntro.vue";
import Features from "./components/FeaturesList.vue";
import Screenshots from "./components/RefBuddyScreenshots.vue";
import ContactForm from "./components/ContactForm.vue";

export default {
    components: {
        Header,
        Introduction,
        Features,
        Screenshots,
        ContactForm,
    },
};
</script>

<style>
@import "@/css/CommonStyles.css";
</style>
