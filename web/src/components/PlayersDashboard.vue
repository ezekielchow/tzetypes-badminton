<script setup lang="ts">

import type { Player } from "@/repositories/clients/private";
import { usePlayerStore } from "@/stores/player-store";
import { useUserStore } from "@/stores/user-store";
import { reactive, ref } from "vue";
import VueTableLite from "vue3-table-lite/ts";

const errorMessage = ref('')
const nameSearch = ref('')

const playerStore = usePlayerStore()
playerStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

const userStore = useUserStore()
userStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

const table = reactive({
    isLoading: false,
    columns: [
        {
            label: "Name",
            field: "name",
            width: "150%",
            sortable: true,
        },
        {
            label: "Actions",
            field: "quick",
            width: "10%",
            display: function (row) {
                return (
                    '<button type="button" data-id="' +
                    row.id +
                    '" class="is-rows-el edit-btn green-button">Edit</button>'
                );
            },
        }
    ],
    rows: [] as Player[],
    totalRecordCount: 0,
    sortable: {
        order: "name",
        sort: "asc",
    },
    pageSize: 10,
    page: 1,
});

const doSearch = async (offset: number, limit: number, order: string, sort: string) => {
    table.isLoading = true;

    let page = 1
    if (offset > 0) {
        page = (offset / limit) + 1
    }

    const res = await userStore.getCurrentUser()

    if (res instanceof Error) {
        errorMessage.value = res.message
        return
    }

    errorMessage.value = ""
    table.page = page

    const listPlayersRes = await playerStore.listPlayers({
        page: table.page,
        pageSize: table.pageSize,
        ownerId: res ? res.id : undefined,
        sortArrangement: `${order}_${sort}`
    })

    if (listPlayersRes instanceof Error) {
        errorMessage.value = listPlayersRes.message
        return
    }

    errorMessage.value = ""
    table.rows = (!!listPlayersRes.players && listPlayersRes.players?.length > 0) ? listPlayersRes.players : []
    table.totalRecordCount = listPlayersRes.pagination?.totalItems ?? 0
};

const tableLoadingFinish = (elements) => {
    table.isLoading = false;
    Array.prototype.forEach.call(elements, function (element) {
        if (element.classList.contains("edit-btn")) {
            element.addEventListener("click", function () {
                console.log(this.dataset.id + " quick-btn click!!");
            });
        }
    });
};


doSearch(0, 0, "name", "asc")

</script>

<template>
    <div class="container">
        <h2><b>Players</b></h2>

        <nav aria-label="Breadcrumb">
            <ol class="breadcrumb">
                <RouterLink to="/dashboard">Dashboard</RouterLink> /
                <RouterLink to="/players">Players</RouterLink>
            </ol>
        </nav>

        <div class="table-wrapper">
            <div v-if="errorMessage" class="error">
                {{ errorMessage }}
            </div>

            <VueTableLite :is-loading="table.isLoading" :columns="table.columns" :rows="table.rows"
                :total="table.totalRecordCount" :sortable="table.sortable" :page="table.page" :pageSize="table.pageSize"
                @do-search="doSearch" @is-finished="tableLoadingFinish" />

        </div>
    </div>
</template>


<style scoped>
.container {
    display: flex;
    flex-direction: column;
}

.breadcrumb {
    margin-top: 1rem;
}
</style>
