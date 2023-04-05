<!-- PrivacyPolicy.vue -->
<template>
    <section class="ref-buddy-page">
        <main class="content">
            <h2>Privacy Policy</h2>
            <p>
                Ref Buddy (the "App") is owned and operated by Jonah Duckworth
                ("us", "we", or "our"). This Privacy Policy describes how we
                collect, use, and disclose your personal information when you
                use our App.
            </p>

            <h3>1. Information Collection</h3>
            <p>
                When you use our App, we may collect the following types of
                information:
            </p>

            <p class="list-item">
                a. Personal Identifiable Information (PII): We may collect your
                name, email address, phone number, and other information that
                can be used to identify you.
            </p>

            <p class="list-item">
                b. Non-Personal Identifiable Information (Non-PII): We may
                collect anonymous usage data, such as device type, operating
                system, and other technical information that does not directly
                identify you.
            </p>

            <p class="list-item">
                c. Location Data: We may collect your device's location data if
                you grant us permission to do so.
            </p>

            <h3>2. How We Use Your Information</h3>
            <p>We use your personal information for the following purposes:</p>

            <p class="list-item">
                a. To provide and maintain the App, including ensuring it
                functions properly and improving its features.
            </p>

            <p class="list-item">
                b. To communicate with you about your use of the App, including
                responding to your inquiries, sending important notices, and
                marketing our products and services.
            </p>

            <p class="list-item">
                c. To analyze usage patterns and improve the App's performance,
                features, and overall user experience.
            </p>

            <p class="list-item">
                d. To detect, prevent, and address technical issues and security
                breaches.
            </p>

            <h3>3. Information Sharing</h3>
            <p>
                We will not sell, rent, or share your personal information with
                third parties except:
            </p>

            <p class="list-item">
                a. When required by law, such as to comply with a subpoena,
                court order, or government request.
            </p>

            <p class="list-item">
                b. To protect our rights, property, or safety, or the rights,
                property, or safety of our users or the public.
            </p>

            <p class="list-item">
                c. With your consent, such as when you opt-in to receive
                marketing materials or participate in a promotion.
            </p>

            <h3>4. Security</h3>
            <p>
                We take the security of your personal information seriously and
                implement reasonable measures to protect it from unauthorized
                access, alteration, disclosure, or destruction. However, no
                method of transmission over the internet or electronic storage
                is 100% secure, and we cannot guarantee absolute security.
            </p>

            <h3>5. Children's Privacy</h3>
            <p>
                Our App is not intended for children under the age of 13, and we
                do not knowingly collect personal information from children
                under 13. If you are a parent or guardian and believe your child
                has provided us with personal information, please contact us
                immediately to have the information removed.
            </p>

            <h3>6. Changes to This Privacy Policy</h3>
            <p>
                We may update our Privacy Policy from time to time. We will
                notify you of any changes by posting the new Privacy Policy on
                our website or within the App. You are advised to review this
                Privacy Policy periodically for any changes.
            </p>

            <h3>7. Contact Us</h3>
            <p>
                If you have any questions about this Privacy Policy, please
                contact us at jonah@refbuddy.ca.
            </p>

            <p>Last updated: April 5, 2023.</p>
        </main>
    </section>
</template>

<script>
export default {
    name: "PrivacyPolicy",
};
</script>

<style>
@import "@/css/CommonStyles.css";

.content {
    flex: 1;
    padding: 16px;
    width: auto;
    max-width: 550px;
    height: 90%;
    overflow-y: scroll;
    margin-left: 10px;
    margin-right: 15px;
    font-size: medium;
}
/* no scroll bar */
::-webkit-scrollbar {
    display: none;
}

/* last updated make font size small */
p:last-child {
    padding-top: 20px;
    font-size: 0.7em;
}

p.list-item {
    margin-left: 1.1em;
    text-indent: -1.1em;
}
</style>
