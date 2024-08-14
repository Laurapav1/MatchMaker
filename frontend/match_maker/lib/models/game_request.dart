import 'package:json_annotation/json_annotation.dart';

part 'game_request.g.dart';

@JsonSerializable()
class GameRequest {
  final int? id;
  final int niveau;
  final String location;
  final String time;
  final String gender;
  final int amount;
  final double price;

  GameRequest(
      {this.id,
      required this.niveau,
      required this.location,
      required this.time,
      required this.gender,
      required this.amount,
      required this.price});

  factory GameRequest.fromJson(Map<String, dynamic> json) => _$GameRequestFromJson(json);

  Map<String, dynamic> toJson() => _$GameRequestToJson(this);
}
