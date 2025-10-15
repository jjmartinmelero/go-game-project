let imgArray = [
    "static/img/hand_rock.png",
    "static/img/hand_paper.png",
    "static/img/hand_scissors.png"
];

function choose(x){
    fetch("play?c="+x)
        .then(res => res.json())
        .then(data => {
            console.log(data)
            if(x == 0) {
                document.getElementById("player_choice").innerHTML = "The player chose rock";
            } else if(x == 1){
                document.getElementById("player_choice").innerHTML = "The player chose paper";
            } else {
                document.getElementById("player_choice").innerHTML = "The player chose scissors";
            }

            document.getElementById("player_score").innerHTML = data.PlayerScore;
            document.getElementById("computer_score").innerHTML = data.ComputerScore;
            document.getElementById("computer_choice").innerHTML = data.ComputerChoice;
            document.getElementById("round_result").innerHTML = data.RoundResult;
            document.getElementById("round_message").innerHTML = data.Message;

            var imgElement = document.getElementById("img_computer");
            imgElement.src = imgArray[data.ComputerChoiceInt];

        })
}