const buttons = document.querySelectorAll('button')

for (const button of buttons) {
    button.addEventListener("click", playSound)
}

function playSound(e) {
    const id = e.target.id
    const url = "/assets/audio/" + id + ".mp3"
    const sound = new Audio(url)
    sound.play()
}