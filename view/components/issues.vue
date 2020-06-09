<template>
  <div>
    <h5 class="card-title">{{ title }}</h5>
    <ul>
      <li v-for="issue in issues" :key="`${title}-issue-${issue.number}`">
        <a target="_blank" :href="issue.html_url"
          >#{{ issue.number }} {{ issue.title }} ({{ issue.state }})</a
        >
        <ul>
          <li
            v-for="comment in issue.comments"
            :key="comment.id"
            class="text-truncate"
          >
            #{{ comment.id }} {{ comment.body }}
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    issues: Array
  }
}
</script>
