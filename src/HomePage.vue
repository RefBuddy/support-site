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
@import url("https://fonts.googleapis.com/css2?family=VT323&display=swap");

body {
    font-family: "VT323", monospace;
    background-color: rgba(255, 145, 0, 0.73);
    font-size: 20px;
}

.ref-buddy-page {
    width: auto;
    max-width: 550px;
    min-width: 300px;
    height: 92%;
    min-height: 450px;
    margin: 0 auto;
    padding: 20px 50px 0 50px;
    background-color: white;
    border-radius: 0;
    border: 3px solid black;
    position: absolute;
    top: calc(50% - 10px);
    left: 50%;
    transform: translate(-50%, -50%);
    box-shadow: 15px 15px 0px 0px black;
}

.header {
    height: 70px;
    width: auto;
    background-color: white;
    border: 1px solid black;
    padding: 16px;
    border-radius: 0;
    text-align: center;
    box-shadow: 5px 5px 0px 0px black;
    position: sticky;
    top: 0;
    left: 0;
    right: 0;
    z-index: 99;
    margin: 0 -16px 16px -16px;
}

.header__image {
    width: 70px;
    height: 70px;
    margin-right: 16px;
    float: left;
    border-radius: 12px;
}

.header__title {
    margin: 0;
    padding-top: 15px;
    margin-right: 15px;
}

.main {
    flex: 1;
    padding: 16px;
    width: auto;
    max-width: 550px;
    height: 78%;
    overflow-y: scroll;
    margin-left: 10px;
    margin-right: 15px;
}

@media screen and (max-height: 760px) {
    .main {
        height: 75%;
    }
}

@media screen and (max-height: 660px) {
    .ref-buddy-page {
        height: 85%;
    }
    .main {
        height: 70%;
    }
}

@media screen and (max-height: 595px) {
    .main {
        height: 65%;
    }
}

@media screen and (max-width: 480px) {
    .ref-buddy-page {
        background-color: transparent;
        border: none;
        box-shadow: none;
        padding: 10px;
        margin: 0;
        height: 100%;
        width: 100%;
        top: 0;
        left: 0;
        transform: none;
    }
    .main {
        height: 80%;
        margin-top: -15px;
    }
    .header {
        margin: 0 15px 15px 0;
    }
}

.main::-webkit-scrollbar {
    display: none;
}

.introduction {
    width: calc(100% - 7px);
    background-color: white;
    border: 1px solid black;
    padding: 16px;
    border-radius: 0;
    box-shadow: 5px 5px 0px 0px black;
    margin: -4px -16px 26px -16px;
}

.features {
    width: calc(100% - 7px);
    background-color: white;
    border: 1px solid black;
    padding: 16px;
    border-radius: 0;
    box-shadow: 5px 5px 0px 0px black;
    margin: -4px -16px 26px -16px;
}

.features__title {
    color: black;
    width: auto;
    max-width: 125px;
    padding: 10px;
    margin: 0 auto;
    text-align: center;
}

.features-list {
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-top: 16px;
}

.features-list__item {
    display: flex;
    gap: 8px;
}

.features-list__icon {
    width: 24px;
    height: 24px;
    object-fit: contain;
}

.features-list__text {
    font-weight: bold;
}

.screenshots {
    width: calc(100% - 7px);
    padding: 16px;
    border-radius: 0;
    text-align: center;
    margin: 0 -16px 26px -16px;
    overflow-x: auto;
    margin-left: -16px;
    margin-right: -16px;
}

.screenshots__title {
    color: black;
    border: 1px solid orange;
    width: auto;
    max-width: 225px;
    padding: 10px;
    margin: 0 auto;
    text-align: center;
}

.screenshots__items {
    display: flex;
    flex-wrap: nowrap;
    overflow-x: auto;
    border: 1px solid rgba(175, 175, 175, 0.236);
    border-radius: 12px;
    box-shadow: 5px 5px 5px 0px rgba(175, 175, 175, 0.236);
}

.screenshots__items > * {
    flex-shrink: 0;
    width: 100%;
    max-width: 100%;
    margin-right: 16px;
}

.screenshots__items > :last-child {
    margin-right: 0;
}

.screenshots__container {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    margin: 0 -16px;
}

.screenshots__arrow {
    background-color: transparent;
    border: none;
    font-size: 32px;
    padding: 8px 16px;
    cursor: pointer;
    color: black;
}

.screenshots__arrow--left {
    margin-right: 16px;
}

.screenshots__arrow--right {
    margin-left: 16px;
}

.footer {
    width: calc(100% - 7px);
    min-height: 245px;
    background-color: white;
    padding: 16px;
    border: 1px solid black;
    border-radius: 0;
    text-align: center;
    box-shadow: 5px 5px 0px 0px black;
    margin: 0 -16px 16px -16px;
}

.footer__success {
    max-height: 70px;
    width: 80%;
    display: flex;
    padding: 16px;
    font-size: 20px;
    background-color: white;
    border: 1px solid black;
    box-shadow: 5px 5px 0px 0px black;
    margin: 0 auto;
}

.footer__form {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 16px;
}

.footer__form > * {
    width: 100%;
    margin-bottom: 16px;
}

.footer__form > textarea {
    height: 100px;
}

.footer__form > button {
    background-color: black;
    color: white;
    border: none;
    padding: 8px 16px;
    cursor: pointer;
}
</style>
