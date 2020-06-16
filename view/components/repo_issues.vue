<template>
  <b-card no-body header="Issues">
    <div v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </div>

    <div v-else>
      <b-list-group>
        <b-list-group-item
          v-for="issue in issues"
          :key="`gitbcket-issue-${issue.number}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="issue.state == 'open'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ issue.number }} {{ issue.title }}
          </span>
          <b-badge v-if="issue.comments.length > 0" variant="primary" pill>{{
            issue.comments.length
          }}</b-badge>
        </b-list-group-item>

        <b-card-body>
          <b-button variant="outline-primary" @click="migrateIssues">
            Migrate
          </b-button>
        </b-card-body>
      </b-list-group>
    </div>
  </b-card>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
  props: ['repo', 'issues', 'loading'],
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken'])
  },
  methods: {
    ...mapActions(['setRepo', 'setProject']),
    async migrateIssues() {
      const res = await this.$axios.$post(
        `/${this.repo.owner.login}/${this.repo.name}/issues`,
        null,
        {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        }
      )

      this.setRepo(res.repo)
      this.setProject(res.project)
    }
  }
}
</script>
