<template>
  <div id="app">
    <section class="introduction">
      <h2 class="component_title">Introduction</h2>
      <p v-if="formattedText(introductionText)" v-html="formattedText(introductionText)"></p>
    </section>

    <section class="introduction">
      <h2 class="component_title">For Officials</h2>
      <p v-if="formattedText(officialsText)" v-html="formattedText(officialsText)"></p>
    </section>

    <section class="introduction">
      <h2 class="component_title">For Assignors</h2>
      <p v-if="formattedText(assignorsText)" v-html="formattedText(assignorsText)"></p>
    </section>

    <section class="introduction">
      <h2 class="component_title">Performance Tracking</h2>
      <p v-if="formattedText(performanceTrackingText)" v-html="formattedText(performanceTrackingText)"></p>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      introductionText: "Ref Buddy is more than just a mobile assigning solution; it's a revolution in managing hockey officiating. \nTailored specifically for hockey officials and assignors, Ref Buddy merges intuitive design with powerful features to transform the way you handle assignments and track performance. With our unique, real-time stat tracking, officials can demonstrate their abilities, growth, and readiness for higher level games. \nThis objective, data-driven approach ensures fair opportunities and aids in the professional development of every official. By focusing not just on assignment, but also on performance tracking, Ref Buddy fosters a culture of continuous improvement and high standards in the world of hockey officiating.",
      officialsText: "With Ref Buddy, officials gain unprecedented control over their careers. Each official has a personalized profile that not only provides access to real-time season stats, but also builds an evolving officiating resume. \nThis detailed record of their experience and accomplishments, backed by objective data, can be a powerful tool when seeking opportunities to officiate higher-level games. The data-driven approach of Ref Buddy promotes continual self-improvement and a fair evaluation of performance. \nWith the ability to review past games, prepare for upcoming ones, and track their progress, officials are equipped with the tools they need to reach their full potential in the world of hockey officiating.",
      assignorsText: "Ref Buddy is not just designed for officials, but also for the assignors who coordinate them. It provides assignors with a comprehensive view of officials' profiles, league schedules, and real-time availability. \nRef Buddy's data-driven approach helps with ensuring the right official is assigned to each game, elevating the standard of officiating and overall game quality.",
      performanceTrackingText: "Ref Buddy elevates the game of hockey by focusing not just on assignment but also on performance tracking. \nWith our cutting-edge mobile solution, every official can track their stats in real-time, fostering a culture of continuous improvement. This objective, data-driven approach offers a transparent and fair system that aids in the professional development of each official. \nWith Ref Buddy, performance tracking isn't just a feature; it's at the core of our mission to enhance the standard of hockey officiating.",
      joinText: "Get ready to streamline your scheduling, enhance performance tracking, and revolutionize the way you manage hockey officiating. \nJoin the Ref Buddy revolution today—the ultimate mobile solution for elite hockey leagues."
    }
  },
  methods: {
    formattedText(text) {
      return text.replace(/\n/g, "<br><br>");
    }
  }
}
</script>