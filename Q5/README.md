# Debugging
## Terjadinya race condition ketika mengubah data dengan go func secara bersamaan maka dari itu dengan cara di-embed ke objek yang memerlukan proses lock-unlock, menjadikan variabel mutex tersebut adalah eksklusif untuk objek tersebut saja.