$(function(){
    const message = $('#msg');

      $('#btn').on('click', function(){

        const content = $('#modal');
        content
          // モーダル開始前の処理
          .on('show.bs.modal', () => {
            console.log('modal open start');
          })
          // モーダル開始後の処理
          .on('shown.bs.modal', () => {
            console.log('modal open complete');
          })
          // モーダル終了前の処理
          .on('hide.bs.modal', () => {
            console.log('modal hidden start');

            // 入力されたメッセージを挿入する
            const str = $('input', content).val();
            if (str.length > 0) {
              message.text(str);
            }
          })
          // モーダル終了後の処理
          .on('hidden.bs.modal', () => {
            console.log('modal hidden complete');

            // 後片付け
            $('input', content).val('');
          })
          .modal({
            backdrop: 'static',
            keyboard: true
          });

        // Close(手動)ボタン
        $('#close', content).on('click', () => {
          content.modal('hide');
        });
      });
});