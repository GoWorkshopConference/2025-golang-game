// ランキング取得・表示機能

// サーバーのベースURL（環境に応じて変更してください）
// HTMLのdata属性から取得するか、直接指定してください
function getServerBaseUrl() {
  const htmlElement = document.documentElement;
  const serverUrl = htmlElement.getAttribute("data-server-url");
  return serverUrl || "http://localhost:8080";
}

export function initRanking() {
  const rankingContainer = document.getElementById("rankingContainer");
  const loadingMessage = document.getElementById("loadingMessage");
  const errorMessage = document.getElementById("errorMessage");
  const rankingList = document.getElementById("rankingList");

  // ランキングを取得
  async function fetchRanking() {
    try {
      // ローディング表示
      loadingMessage.style.display = "block";
      errorMessage.style.display = "none";
      rankingList.innerHTML = "";

      const response = await fetch(`${getServerBaseUrl()}/ranking`);
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      
      // ローディング非表示
      loadingMessage.style.display = "none";

      if (data.rankings && data.rankings.length > 0) {
        // ランキングを表示
        data.rankings.forEach((score, index) => {
          const rank = index + 1;
          const listItem = document.createElement("div");
          listItem.className = "ranking-item";
          
          const rankElement = document.createElement("span");
          rankElement.className = "rank";
          rankElement.textContent = `${rank}位`;

          const nameElement = document.createElement("span");
          nameElement.className = "username";
          nameElement.textContent = score.username || "名無し";

          const scoreElement = document.createElement("span");
          scoreElement.className = "score";
          scoreElement.textContent = `${score.score.toLocaleString()}点`;

          const dateElement = document.createElement("span");
          dateElement.className = "date";
          const date = new Date(score.created_at?.seconds * 1000 || score.created_at);
          dateElement.textContent = formatDate(date);

          listItem.appendChild(rankElement);
          listItem.appendChild(nameElement);
          listItem.appendChild(scoreElement);
          listItem.appendChild(dateElement);
          
          rankingList.appendChild(listItem);
        });
      } else {
        rankingList.innerHTML = '<div class="no-ranking">ランキングデータがありません</div>';
      }
    } catch (error) {
      console.error("ランキングの取得に失敗しました:", error);
      loadingMessage.style.display = "none";
      errorMessage.style.display = "block";
      errorMessage.textContent = `ランキングの取得に失敗しました: ${error.message}`;
    }
  }

  // 日付をフォーマット
  function formatDate(date) {
    if (!date || isNaN(date.getTime())) {
      return "-";
    }
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");
    return `${year}/${month}/${day}`;
  }

  // 初回読み込み時にランキングを取得
  fetchRanking();

  // リフレッシュボタンがあれば、クリック時に再取得
  const refreshButton = document.getElementById("refreshRankingButton");
  if (refreshButton) {
    refreshButton.addEventListener("click", fetchRanking);
  }

  // 定期的に更新する場合は、以下のコメントを外してください
  // setInterval(fetchRanking, 30000); // 30秒ごとに更新
}

