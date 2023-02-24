
var intervalRejoin = window.setInterval(rejoin, 1000);

function rejoin(interval){
    var bth = document.querySelector('button.ts-btn.ts-btn-fluent.ts-btn-fluent-primary')
    if (bth!==null){
        console.log("Rejoining")
        bth.click()
        }
    else{
        console.log("No need to rejoin")
    }
}
