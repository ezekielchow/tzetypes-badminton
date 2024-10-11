<script setup lang="ts">
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
            <table class="responsive-table">
                <thead>
                    <tr>
                        <th>Header 1</th>
                        <th>Header 2</th>
                        <th>Header 3</th>
                        <th>Header 4</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Data 1</td>
                        <td>Data 2</td>
                        <td>Data 3</td>
                        <td>Data 4</td>
                    </tr>
                    <tr>
                        <td>Data 5</td>
                        <td>Data 6</td>
                        <td>Data 7</td>
                        <td>Data 8</td>
                    </tr>
                    <tr>
                        <td>Data 9</td>
                        <td>Data 10</td>
                        <td>Data 11</td>
                        <td>Data 12</td>
                    </tr>
                </tbody>
            </table>
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
