<template>
    <div class="ref-buddy-page">
        <header class="header">
            <img
                src="@/assets/Ref_Buddy.jpg"
                class="header__image"
                alt="Ref Buddy logo"
            />
            <h1 class="header__title">Ref Buddy</h1>
        </header>
        <main class="main">
            <section class="introduction">
                <h2 class="introduction__title">What is Ref Buddy?</h2>
                <p class="introduction__description">
                    Ref Buddy is a mobile scheduling system used to assign
                    on-ice officials to hockey games.
                </p>
                <p class="introduction__description">
                    Each official gets their own profile where they can see
                    their season stats, view and update availability, see
                    upcoming games, and view previous games.
                </p>
                <p class="introduction__description">
                    Ref Buddy allows for assignors to quickly and seemlessly add
                    officials to games, as well as easily view their stats,
                    previous games, and availability.
                </p>
            </section>
            <section class="features">
                <h2 class="features__title">Features</h2>
                <ul class="features__list">
                    <li>Schedule officiating for hockey leagues</li>
                    <li>
                        AI automatically collects and imports league schedule
                    </li>
                    <li>Individual official profiles</li>
                    <li>AI generated stats for each official</li>
                    <li>
                        Quick access to all assigned games, along with their box
                        scores
                    </li>
                    <li>
                        When an official gets assigned a game, they receive a
                        notification from Ref Buddy on their phone
                    </li>
                    <li>More to come!</li>
                </ul>
            </section>
            <section class="screenshots">
                <h2 class="screenshots__title">Screenshots</h2>
                <div class="screenshots__items">
                    <img src="screenshot-1.png" alt="Screenshot 1" />
                    <img src="screenshot-2.png" alt="Screenshot 2" />
                    <img src="screenshot-3.png" alt="Screenshot 3" />
                </div>
            </section>
            <footer class="footer">
                <h2 class="footer__title">Contact Us</h2>
                <form class="footer__form">
                    <input type="text" placeholder="Name" />
                    <input type="email" placeholder="Email" />
                    <textarea placeholder="Message"></textarea>
                    <button type="submit">Submit</button>
                </form>
            </footer>
        </main>
    </div>
</template>

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
    height: 80%;
    margin: 0 auto;
    padding: 40px;
    background-color: white;
    border-radius: 0;
    border: 3px solid black;
    position: absolute;
    top: 50%;
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
    height: 80%;
    overflow-y: scroll;
    margin-left: 10px;
}

/* Hide the scrollbar */
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

.introduction__title {
    color: black;
    border: 1px solid orange;
    width: auto;
    max-width: 225px;
    padding: 10px;
    margin: 0 auto;
    text-align: center;
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
    border: 1px solid orange;
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
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    background-color: white;
    border: 1px solid black;
    padding: 16px;
    border-radius: 0;
    text-align: center;
    box-shadow: 5px 5px 0px 0px black;
    margin: 0 -16px 26px -16px;
}

.screenshots__title {
    color: black;
    text-align: center;
    margin-bottom: 16px;
}

.screenshots__items {
    display: flex;
}

.screenshots__items > * {
    width: 30%;
    margin: 0 16px 16px 0;
}

.footer {
    width: calc(100% - 7px);
    background-color: white;
    padding: 16px;
    border: 1px solid black;
    border-radius: 0;
    text-align: center;
    box-shadow: 5px 5px 0px 0px black;
    margin: 0 -16px 16px -16px;
}

.footer__title {
    color: black;
    margin-bottom: 16px;
}

.footer__form {
    display: flex;
    flex-direction: column;
    align-items: center;
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
