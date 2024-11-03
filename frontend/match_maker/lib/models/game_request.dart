import 'package:json_annotation/json_annotation.dart';

part 'game_request.g.dart';

@JsonSerializable()
class GameRequest {
  @JsonKey(name: 'ID')
  final int? id;

  @JsonKey(name: 'UserEmail')
  final String userEmail;

  @JsonKey(name: 'Niveau')
  final int niveau;

  @JsonKey(name: 'Location')
  final String location;

  @JsonKey(name: 'Time')
  final DateTime time;

  @JsonKey(name: 'Gender')
  final String gender;

  @JsonKey(name: 'Amount')
  final int amount;

  @JsonKey(name: 'Price')
  final double price;

  GameRequest({
    this.id,
    required this.userEmail,
    required this.niveau,
    required this.location,
    required this.time,
    required this.gender,
    required this.amount,
    required this.price,
  });

  factory GameRequest.fromJson(Map<String, dynamic> json) =>
      _$GameRequestFromJson(json);

  Map<String, dynamic> toJson() => _$GameRequestToJson(this);
}
