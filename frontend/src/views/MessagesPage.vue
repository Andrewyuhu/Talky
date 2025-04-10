<script lang="ts" setup>
import MainLayout from "../layouts/MainLayout.vue";
import ActiveChatPane from "../components/ActiveChatPane.vue";
import ChatsPane from "../components/ChatsPane.vue";
</script>

<template>
  <MainLayout>
    <div class="grid grid-row-[1fr] grid-cols-3 h-full min-h-0">
      <ChatsPane></ChatsPane>
      <ActiveChatPane class="col-span-2"></ActiveChatPane>
    </div>
  </MainLayout>
</template>
