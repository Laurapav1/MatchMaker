// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'game_request.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

GameRequest _$GameRequestFromJson(Map<String, dynamic> json) => GameRequest(
      id: (json['ID'] as num?)?.toInt(),
      userEmail: json['UserEmail'] as String,
      niveau: (json['Niveau'] as num).toInt(),
      location: json['Location'] as String,
      time: DateTime.parse(json['Time'] as String),
      gender: json['Gender'] as String,
      amount: (json['Amount'] as num).toInt(),
      price: (json['Price'] as num).toDouble(),
    );

Map<String, dynamic> _$GameRequestToJson(GameRequest instance) =>
    <String, dynamic>{
      'ID': instance.id,
      'UserEmail': instance.userEmail,
      'Niveau': instance.niveau,
      'Location': instance.location,
      'Time': instance.time.toIso8601String(),
      'Gender': instance.gender,
      'Amount': instance.amount,
      'Price': instance.price,
    };
