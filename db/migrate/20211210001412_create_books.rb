class CreateBooks < ActiveRecord::Migration[5.2]
  def up
    create_table :books do |t|
      t.integer :publisher_id
      t.integer :category_id
      t.string :name
      t.string :author
      t.integer :year_of_publication
      t.integer :stock
      t.integer :price
      t.timestamps
    end
  end

  def down 
    drop_table :books
  end
end
