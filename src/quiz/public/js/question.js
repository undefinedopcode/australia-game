var gameState = {}

const correctNeeded = 5;
const roundSize = 10;

function newGame() {
    gameState = {
        score: 0,
        question: null,
        selectedAnswer: -1,
        questions: 0,
        maxQuestions: roundSize,
        correctThisRound: 0,
        round: 1,
    };
    fetchQuestion();
}

function newRound() {
    gameState = {
        score: gameState.score,
        question: null,
        selectedAnswer: -1,
        questions: 0,
        maxQuestions: roundSize,
        correctThisRound: 0,
        round: gameState.round + 1,
    };
    fetchQuestion();
}

function renderScore() {
    $("#score-value").html( gameState.score );
    $("#question-number").html( gameState.questions );
    $("#question-count").html( gameState.maxQuestions );
    $("#round-number").html( gameState.round );
}

function renderQuestion() {

    updateMap( 
        gameState.question["region"]["lat"], 
        gameState.question["region"]["lon"],
        gameState.question["region"]["nelat"], 
        gameState.question["region"]["nelon"],
        gameState.question["region"]["swlat"], 
        gameState.question["region"]["swlon"],
    );

    $('#question-text').html(gameState.question["text"]);
    $('#text-a').html(gameState.question["answers"][0]["text"]);
    $('#text-b').html(gameState.question["answers"][1]["text"]);
    $('#text-c').html(gameState.question["answers"][2]["text"]);
    $('#text-d').html(gameState.question["answers"][3]["text"]);

    clearHighlights();
    gameState.selectedAnswer = -1

    renderScore();

}

function clearHighlights() {
    $('#answer-a').removeClass( 'select' );
    $('#answer-b').removeClass( 'select' );
    $('#answer-c').removeClass( 'select' );
    $('#answer-d').removeClass( 'select' );
    $('#answer-a').removeClass( 'correct' );
    $('#answer-b').removeClass( 'correct' );
    $('#answer-c').removeClass( 'correct' );
    $('#answer-d').removeClass( 'correct' );
    $('#answer-a').removeClass( 'incorrect' );
    $('#answer-b').removeClass( 'incorrect' );
    $('#answer-c').removeClass( 'incorrect' );
    $('#answer-d').removeClass( 'incorrect' );
}

function highlightAnswer(answer, c) {
    switch (answer) {
        case 0: $('#answer-a').addClass( c );
            break;
        case 1: $('#answer-b').addClass( c );
            break;
        case 2: $('#answer-c').addClass( c );
            break;
        case 3: $('#answer-d').addClass( c );
            break;
    }
}

function selectAnswer(answer) {
    clearHighlights();

    highlightAnswer(answer, 'select')

    console.log("selected answer "+answer);

    gameState.selectedAnswer = Number(answer);

    setTimeout( function() {
        validateAnswer();
    }, 1000 );
}

function validateAnswer() {

    if (gameState.selectedAnswer === -1) {
        return;
    }

    if (gameState.selectedAnswer === gameState.question["correct"]) {
        // yay - make it green!
        clearHighlights();
        highlightAnswer( gameState.selectedAnswer, "correct" );
        gameState.score += 1;
        gameState.correctThisRound += 1;
    } else {
        // boo - make it red
        clearHighlights();
        highlightAnswer( gameState.selectedAnswer, "incorrect" );
        highlightAnswer( gameState.question['correct'], "correct" );
    }

    renderScore();

    setTimeout(
        function() {
            if (gameState.questions === roundSize) {
                // next round?
                if (gameState.correctThisRound < correctNeeded) {
                    gameOver();
                } else {
                    newRound();
                }
            } else {
                fetchQuestion();
            }
        }, 
        2000,
    )

}

function gameOver() {
    console.log("game ended");
}

function fetchQuestion() {

    fetch('/api/questions/population')
    .then(function(response) {
      return response.json();
    })
    .then(function(myJson) {
      console.log(JSON.stringify(myJson));
      gameState.question = myJson;
      gameState.questions += 1;
      renderQuestion();
    });

}

$( document ).ready(function() {
    $("#answer-a").click(function(e){     //function_td
        selectAnswer(0);
        e.stopPropagation();
    });
    $("#answer-b").click(function(e){     //function_td
        selectAnswer(1);
        e.stopPropagation();
    });
    $("#answer-c").click(function(e){     //function_td
        selectAnswer(2);
        e.stopPropagation();
    });
    $("#answer-d").click(function(e){     //function_td
        selectAnswer(3);
        e.stopPropagation();
    });
    newGame();
});