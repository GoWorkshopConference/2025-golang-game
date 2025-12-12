// プレイヤー名の入力・保存機能

export function initPlayerName(iframe) {
  const playerNameInput = document.getElementById("playerNameInput");
  const saveNameButton = document.getElementById("saveNameButton");

  // ページ読み込み時に保存されている名前を表示
  const savedName = localStorage.getItem("playerName");
  if (savedName) {
    playerNameInput.value = savedName;
  }

  // 保存ボタンのクリックイベント
  saveNameButton.addEventListener("click", () => {
    const playerName = playerNameInput.value.trim();
    if (playerName) {
      localStorage.setItem("playerName", playerName);
      // iframeに通知を送信
      iframe.contentWindow.postMessage(
        {
          type: "localStorage",
          data: { playerName: playerName },
        },
        "*",
      );
      alert("名前を保存しました！");
    } else {
      alert("名前を入力してください");
    }
  });

  // Enterキーでも保存できるようにする
  playerNameInput.addEventListener("keypress", (e) => {
    if (e.key === "Enter") {
      saveNameButton.click();
    }
  });
}

