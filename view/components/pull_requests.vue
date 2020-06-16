<template>
  <b-card no-body header="Pull Requests">
    <div v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </div>

    <div v-else>
      <b-list-group>
        <b-list-group-item
          v-for="pull in pulls"
          :key="`pull-${pull.number}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="pull.state == 'open'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ pull.number }} {{ pull.title }}
          </span>
          <b-badge v-if="pull.comments.length > 0" variant="primary" pill>{{
            pull.comments.length
          }}</b-badge>
        </b-list-group-item>

        <b-card-body>
          <b-button variant="outline-primary" @click="migratePulls">
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
  props: ['repo', 'pulls', 'loading'],
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken'])
  },
  methods: {
    ...mapActions(['setRepo', 'setProject']),
    async migratePulls() {
      const res = await this.$axios.$post(
        `/${this.repo.owner.login}/${this.repo.name}/pulls`,
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

<style></style>
