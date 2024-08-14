// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'game_request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

GameRequest _$GameRequestFromJson(Map<String, dynamic> json) => GameRequest(
      id: (json['id'] as num).toInt(),
      niveau: (json['niveau'] as num).toInt(),
      location: json['location'] as String,
      time: json['time'] as String,
      gender: json['gender'] as String,
      amount: (json['amount'] as num).toInt(),
      price: (json['price'] as num).toDouble(),
    );

Map<String, dynamic> _$GameRequestToJson(GameRequest instance) =>
    <String, dynamic>{
      'id': instance.id,
      'niveau': instance.niveau,
      'location': instance.location,
      'time': instance.time,
      'gender': instance.gender,
      'amount': instance.amount,
      'price': instance.price,
    };
