<template>
  <b-card no-body header="Merge Requests">
    <div v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </div>

    <div v-else>
      <b-list-group flush>
        <b-list-group-item
          v-for="merge in merges"
          :key="`merge-${merge.iid}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="merge.state == 'opened'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ merge.iid }} {{ merge.title }}
          </span>
          <b-badge
            v-if="merge.comments && merge.comments.length > 0"
            variant="primary"
            pill
            >{{ merge.comments.length }}</b-badge
          >
        </b-list-group-item>
      </b-list-group>
    </div>
  </b-card>
</template>

<script>
export default {
  props: ['merges', 'loading']
}
</script>

<style></style>
