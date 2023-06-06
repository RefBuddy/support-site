<template>
    <div id="app">
        <section class="introduction">
            <h2 class="component_title">Introduction</h2>
            <p v-if="formattedIntroductionText" v-html="formattedIntroductionText"></p>
            <p v-else>{{ introductionText }}</p>
        </section>
        
        <section class="introduction">
            <h2 class="component_title">For Officials</h2>
            <p v-if="formattedOfficialsText" v-html="formattedOfficialsText"></p>
            <p v-else>{{ officialsText }}</p>
        </section>
        
        <section class="introduction">
            <h2 class="component_title">For Assignors</h2>
            <p>{{ assignorsText }}</p>
        </section>
        
        <section class="introduction">
            <h2 class="component_title">Performance Tracking</h2>
            <p>{{ performanceTrackingText }}</p>
        </section>
        
        <section class="introduction">
            <h2 class="component_title">Join the Revolution</h2>
            <p>{{ joinText }}</p>
        </section>
    </div>
</template>
  
<script>
  export default {
    data() {
      return {
        introductionText: "Ref Buddy is more than just a mobile assigning solution; it's a revolution in managing hockey officiating. \nTailored specifically for hockey officials and assignors, Ref Buddy merges intuitive design with powerful features to transform the way you handle assignments and track performance. With our unique, real-time stat tracking, officials can demonstrate their abilities, growth, and readiness for higher level games. \nThis objective, data-driven approach ensures fair opportunities and aids in the professional development of every official. By focusing not just on assignment, but also on performance tracking, Ref Buddy fosters a culture of continuous improvement and high standards in the world of hockey officiating.",
        officialsText: "With Ref Buddy, officials gain unprecedented control over their careers. Each official has a personalized profile that not only provides access to real-time season stats, but also builds an evolving officiating resume. \nThis detailed record of their experience and accomplishments, backed by objective data, can be a powerful tool when seeking opportunities to officiate higher-level games. The data-driven approach of Ref Buddy promotes continual self-improvement and a fair evaluation of performance. \nWith the ability to review past games, prepare for upcoming ones, and track their progress, officials are equipped with the tools they need to reach their full potential in the world of hockey officiating.",
        assignorsText: "For assignors, Ref Buddy is the ultimate tool for efficient scheduling. It provides comprehensive access to Ref profiles, league schedules, and real-time availability updates, all in one convenient place. With Ref Buddy, assigning the right official to the right game becomes a smooth and streamlined process.",
        performanceTrackingText: "But Ref Buddy offers more than just assignment management—it's also a powerful performance tracking tool. Our unique platform offers real-time stat tracking for officials, a feature that sets us apart from other assigning services. This allows leagues and officials to monitor performance closely, fostering a culture of continuous improvement and high-quality officiating.",
        joinText: "Get ready to streamline your scheduling, enhance performance tracking, and revolutionize the way you manage hockey officiating. Join the Ref Buddy revolution today—the ultimate mobile solution for elite hockey leagues."
    }
  },
  computed: {
    formattedIntroductionText() {
      return this.introductionText.replace(/\n/g, "<br><br>");
    },
    formattedOfficialsText() {
      return this.officialsText.replace(/\n/g, "<br><br>");
    }
  }
}
</script>
  
<style scoped>
  section {
    margin-bottom: 2em;
  }
  
  h1, h2 {
    color: #333;
  }
  
  p {
    color: #666;
  }
</style>
  
