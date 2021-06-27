/* vim: set ts=4 sw=4: */

/*
handler はControllersから処理を委譲され、各コマンドの実行の責任を担います。
主な役割としてUseCaseへの橋渡しとして動作し、コマンド毎で受付可能なバリデーションの
責務を負います。

CleanArchitectureのInterface Adaptersにおける振る舞いを実装しています。
*/
package handler
