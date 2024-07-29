class GameRequest {
  final int id;
  final int niveau;
  final String location;
  final String time;
  final String gender;
  final int amount;
  final double price;

  const GameRequest({
    required this.id,
    required this.niveau,
    required this.location,
    required this.time,
    required this.gender,
    required this.amount,
    required this.price
  });

  factory GameRequest.fromJson(Map<String, dynamic> json) {
    return switch (json) {
      {
        'id': int id,
        'niveau': int niveau,
        'location': String location,
        'time': String time,
        'gender': String gender,
        'amount': int amount,
        'price': double price,
      } =>
        GameRequest(
          id: id,
          niveau: niveau,
          location: location,
          time: time,
          gender: gender,
          amount: amount,
          price: price
        ),
      _ => throw const FormatException('Failed to load match requests.')
    };
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['id'] = id;
    data['niveau'] = niveau;
    data['location'] = location;
    data['time'] = time;
    data['gender'] = gender;
    data['amount'] = amount;
    data['price'] = price;
    return data;
  }
}