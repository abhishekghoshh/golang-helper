let socket;

// generate a random ID for typing messages
let typing_id = Math.floor(Math.random() * 1000000);

// Function to generate a random Materialize color class
function getRandomColorClass() {
  const colors = [
    "red", "pink", "purple", "deep-purple", 
    "indigo", "blue", "light-blue", "cyan", 
    "teal", "green", "light-green", "lime", 
    "yellow", "amber", "orange", "deep-orange", 
    "brown", "grey", "blue-grey"
  ];
  const randomIndex = Math.floor(Math.random() * colors.length);
  return colors[randomIndex];
}

// Object to store sender-to-color mapping
const senderColors = {};

$(document).ready(function () {
  // Connect button
  $("#connect-btn").click(function () {
    if (socket) {
      M.toast({ html: "Already connected!" });
      return;
    }
    
    // Open WebSocket connection
    socket = new WebSocket("/chat");

    socket.onopen = function () {
      M.toast({ html: "Connected to the chat!" });
    };

    socket.onmessage = function (event) {
      const message = event.data;

      if (message.includes(":")) {
        // Split the message at the first colon
        const [sender, ...messageParts] = message.split(":");
        const actualMessage = messageParts.join(":").trim();

        if (actualMessage.startsWith("INTERMEDIATE_MSG+")) {
          // Handle intermediate typing message
          const typingMessage = actualMessage.replace("INTERMEDIATE_MSG+", "").trim();
          const [typingId, ...typingParts] = typingMessage.split("+");
          const typingContent = typingParts.join("+").trim();
          // if there is li with "typing-" + typing_id, then update its content, otherwise create a new li
          if ($(`#typing-${typingId}`).length) {
            $(`#typing-${typingId} .typing-message`).html(`<em>${sender} is typing: ${typingContent}</em>`);
          } else {
           // Optionally, you can display the typing message in a different style
            const $typingMsg = $(`
              <li class="collection-item typing-message-item" id="typing-${typingId}">
                <div class="typing-message-content">
                  <span class="typing-message"><em>${sender} is typing: ${typingContent}</em></span>
                </div>
              </li>
            `);
            $("#chat-box").append($typingMsg);
          }
          setTimeout(() => {
            $(`#typing-${typingId}`).remove(); // Remove the typing message after 1 second
            console.log(`Removed typing message for ${sender} with ID: ${typingId}`);
          }, 5000);
          return; // Skip further processing for intermediate messages
        }

        // Assign a color to the sender if not already assigned
        if (!senderColors[sender]) {
          senderColors[sender] = getRandomColorClass();
        }

        const colorClass = senderColors[sender];

        // Append message with sender info and profile symbol
        $("#chat-box").append(`
        <li class="collection-item avatar custom-message-item">
            <i class="material-icons circle ${colorClass}"></i>
            <div class="message-content">
                <span class="title sender-name"><strong>${sender}</strong></span>
                <p class="message-text">${actualMessage}</p>
            </div>
        </li>
        `);
      } else {
        // Handle messages without colons as highlighted system messages
        $("#chat-box").append(`
        <li class="collection-item highlighted-message-item">
            <div class="highlighted-message-content">
                <span class="highlighted-message">${message}</span>
            </div>
        </li>
        `);
      }

      // Scroll to the bottom of the chat box
      const chatBox = document.getElementById("chat-box");
      chatBox.scrollTop = chatBox.scrollHeight;
    };

    socket.onerror = function () {
      M.toast({ html: "Error connecting to the server!" });
    };

    socket.onclose = function () {
      M.toast({ html: "Connection closed." });
      socket = null;
    };
  });

  // Send message on Enter
  $("#message-input").keypress(function (e) {
    if (e.which == 13 && socket && socket.readyState === WebSocket.OPEN) {
      const message = $(this).val();
      if (message.trim() !== "") {
        socket.send(message);
        $(this).val(""); // Clear input
        // if there is an intermediate typing message, remove it immediately after sending the message
        $(`.typing-message-item`).remove();
        typing_id = Math.floor(Math.random() * 1000000); // Generate a new typing ID for the next message
      }
    }
  });
  // Send message on typing
  $("#message-input").on("input", function () {
    if (socket && socket.readyState === WebSocket.OPEN) {
      const typingMessage = "INTERMEDIATE_MSG+"+typing_id+"+"+$(this).val();
      socket.send(typingMessage);
    }
  });
});
